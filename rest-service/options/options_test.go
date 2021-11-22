package options

import (
	"testing"
)

func TestOptions(t *testing.T) {
	opt, err := InitOptions()
	if err != nil {
		t.Error(err)
	}
	if opt.Port != "8090" {
		t.Errorf("Expected port to be 8090, got %v", opt.Port)
	}

	if !opt.Debug {
		t.Errorf("Expected debug to be true, got %t", opt.Debug)
	}
}
