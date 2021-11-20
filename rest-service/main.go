package main

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/stackpath/backend-developer-tests/rest-service/global"
	coreserver "github.com/stackpath/backend-developer-tests/rest-service/pkg/core"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()
	ctx, cancel := context.WithCancel(context.Background())
	// Log every logrus message
	if global.Options.Debug {
		log.SetLevel(log.DebugLevel)
	}

	// Initialize Rest APIs and start graceful server
	err := coreserver.CoreServer(ctx, cancel)
	if err != nil {
		//send signal to all channels to calm down we found an error
		log.Errorf("Backend server went crazy: %v", err.Error())
		// send OS signal and shut it all down
		os.Exit(1)
	}
	log.Infof("Backend intialization completed")
}
