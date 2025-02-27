// Package global contains global variables to be used across the backend
package global

import (
	"sync"

	"context"

	"github.com/stackpath/backend-developer-tests/rest-service/options"

	log "github.com/sirupsen/logrus"
)

// some global variables commonly used
var (
	Options         *options.Options
	UnitTest        bool
	Ctx             context.Context
	WaitGroupServer sync.WaitGroup
)

// initializes global package to read environment variables as needed
func init() {
	options, err := options.InitOptions()
	if err != nil {
		log.Fatal("Options init errored: ", err.Error())
	}
	Options = options
}
