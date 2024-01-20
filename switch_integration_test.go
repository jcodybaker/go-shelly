package shelly_test

import (
	"testing"

	"github.com/jcodybaker/go-shelly"
)

func TestSwitchGetConfig(t *testing.T) {
	req := &shelly.SwitchGetConfigRequest{
		ID: 0,
	}
	resp := req.NewTypedResponse()
	GetCallWithVerify(t, req, resp)
}

func TestSwitchGetStatus(t *testing.T) {
	req := &shelly.SwitchGetStatusRequest{
		ID: 0,
	}
	resp := req.NewTypedResponse()
	GetCallWithVerify(t, req, resp)
}
