package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

type CloudSetConfigRequest struct {
	Config CloudConfig `json:"config"`
}

func (r *CloudSetConfigRequest) Method() string {
	return "Cloud.SetConfig"
}

func (r *CloudSetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *CloudSetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CloudSetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*SetConfigResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type CloudConfig struct {
	// Enable is true if cloud connection is enabled, false otherwise
	Enable bool `json:"enable"`

	// Server is the name of the server to which the device is connected (optional).
	Server *string `json:"server"`
}

type CloudGetConfigRequest struct{}

func (r *CloudGetConfigRequest) Method() string {
	return "Cloud.GetConfig"
}

func (r *CloudGetConfigRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *CloudGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CloudGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*RPCEmptyResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type CloudStatus struct {
	Connected bool `json:"connected"`
}

type CloudGetStatusRequest struct{}

func (r *CloudGetStatusRequest) Method() string {
	return "Cloud.GetStatus"
}

func (r *CloudGetStatusRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *CloudGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CloudGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*RPCEmptyResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}
