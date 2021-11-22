package core

import (
	"context"
	"testing"

	"github.com/stackpath/backend-developer-tests/rest-service/global"
)

// Simple test to check if the global package is working if serve start successfully
func TestController(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// start the backend controller and wait for it to finish booting up
	CoreServer(ctx, cancel)
	global.WaitGroupServer.Done()
	cancel()
}
