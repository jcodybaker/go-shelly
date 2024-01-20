package shelly_test

import (
	"testing"

	"github.com/jcodybaker/go-shelly"
)

func TestWifiGetConfig(t *testing.T) {
	req := &shelly.WifiGetConfigRequest{}
	resp := req.NewTypedResponse()
	GetCallWithVerify(t, req, resp)
}

func TestWifiGetStatus(t *testing.T) {
	req := &shelly.WifiGetStatusRequest{}
	resp := req.NewTypedResponse()
	GetCallWithVerify(t, req, resp)
}
