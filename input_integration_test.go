package shelly_test

import (
	"testing"

	"github.com/jcodybaker/go-shelly"
)

func TestInputGetConfig(t *testing.T) {
	req := &shelly.InputGetConfigRequest{
		ID: 0,
	}
	resp := req.NewTypedResponse()
	GetCallWithVerify(t, req, resp)
}

func TestInputGetStatus(t *testing.T) {
	req := &shelly.InputGetStatusRequest{
		ID: 0,
	}
	resp := req.NewTypedResponse()
	GetCallWithVerify(t, req, resp)
}
