// Package trigger entry point of go service
package core

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/stackpath/backend-developer-tests/rest-service/pkg/core/controller"
	"github.com/stackpath/backend-developer-tests/rest-service/global"
)

// CoreServer ....CoreServer
func CoreServer(ctx context.Context, cancel context.CancelFunc) error {
	// create a new context and save it in global
	port := 8090
	log.Infof("starting backend service.")
	portFromEnv := global.Options.Port;
	var err error
	if portFromEnv != "" {
		port, err = strconv.Atoi(portFromEnv)
		if err != nil {
			port = 8090
		}
	}
	global.Ctx = ctx

	// setup cancel signal for graceful shutdown of serve\
	go monitorSystem(cancel)
	bootUPErrors := make(chan error, 1)

	controller.BackendController(ctx, port, bootUPErrors)
	//........................................................................................
	// Block the server until server encounter any errors
	err = <-bootUPErrors
	if err != nil {
		log.Errorf("There is an issue starting backend server : %v", err.Error())
		global.WaitGroupServer.Wait()
		return err
	}
	return nil
}

func monitorSystem(cancel context.CancelFunc) {
	holdSignal := make(chan os.Signal, 1)
	signal.Notify(holdSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	// if system throw any termination stuff let channel handle it and cancel
	<-holdSignal
	cancel()
}
