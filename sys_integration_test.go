package shelly_test

import (
	"testing"

	"github.com/jcodybaker/go-shelly"
)

func TestSysGetConfig(t *testing.T) {
	req := &shelly.SysGetConfigRequest{}
	resp := req.NewTypedResponse()
	GetCallWithVerify(t, req, resp)
}

func TestSysGetStatus(t *testing.T) {
	req := &shelly.SysGetStatusRequest{}
	resp := req.NewTypedResponse()
	GetCallWithVerify(t, req, resp)
}
