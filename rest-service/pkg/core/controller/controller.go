package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stackpath/backend-developer-tests/rest-service/constants"
	"github.com/stackpath/backend-developer-tests/rest-service/global"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/core/router"
	graceful "gopkg.in/tylerb/graceful.v1" // see: https://github.com/tylerb/graceful
)

type maxPayloadHandler struct {
	handler http.Handler
	size    int64
}

// ServeHTTP uses MaxByteReader to limit the size of the input
func (handler *maxPayloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, handler.size)
	handler.handler.ServeHTTP(w, r)
}

// BackendController ....
func BackendController(ctx context.Context, port int, errorChannel chan error) {
	log.Infof("Initializing router and endpoints.")
	backendRouter, err := router.BackendRouter()
	if err != nil {
		errorChannel <- err
		return
	}
	httpAddress := fmt.Sprintf(":%d", port)
	// set maximum payload size for incoming requests

	// init maxHandler to limit size of request input
	var httpHandler http.Handler
	if global.Options.MaxPayloadSize > 0 {
		httpHandler = &maxPayloadHandler{handler: backendRouter, size: global.Options.MaxPayloadSize}
	} else {
		httpHandler = backendRouter
	}
	global.WaitGroupServer.Add(1)
	go serverHTTPRoutes(ctx, httpAddress, httpHandler, errorChannel)
}

func serverHTTPRoutes(ctx context.Context, httpAddress string, handler http.Handler, errorChannel <-chan error) {
	defer global.WaitGroupServer.Done()
	// init graceful server
	serverGrace := &graceful.Server{
		Timeout: 10 * time.Second,
		//BeforeShutdown:    beforeShutDown,
		ShutdownInitiated: shutDownBackend,
		Server: &http.Server{
			Addr:           httpAddress,
			Handler:        handler,
			MaxHeaderBytes: global.Options.MaxHeaderSize,
			ReadTimeout:    time.Duration(constants.MAX_READ_TIMEOUT) * time.Second,
			WriteTimeout:   time.Duration(constants.MAX_WRITE_TIMEOUT) * time.Second,
		},
	}
	stopChannel := serverGrace.StopChan()
	err := serverGrace.ListenAndServe()
	if err != nil {
		log.Fatalf("BackendController: Failed to start server : %s", err.Error())
	}
	log.Infof("Backend is serving the routes.")
	for {
		// wait for the server to stop or be canceled
		select {
		case <-stopChannel:
			log.Infof("BackendController: Server shutdown at %s", time.Now())
			os.Exit(1)
			return
		case <-ctx.Done():
			log.Infof("BackendController: context done is called %s", time.Now())
			serverGrace.Stop(time.Second * 2)
			// call sigterm and set exit signal
			os.Exit(1)
		}
	}
}

func shutDownBackend() {
	log.Infof("BackendController: Shutting down server at %s", time.Now())
}
