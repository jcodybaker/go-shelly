package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

type BLEStatus struct{}

type BLEGetStatusRequest struct{}

func (r *BLEGetStatusRequest) Method() string {
	return "BLE.GetStatus"
}

func (r *BLEGetStatusRequest) NewTypedResponse() *BLEStatus {
	return &BLEStatus{}
}

func (r *BLEGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BLEGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*BLEStatus,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

type BLEConfig struct {
	// Enable is true if bluetooth is enabled, false otherwise.
	Enable *bool `json:"enable,omitempty"`

	// RPC is the configuration of the rpc service.
	RPC *BLERPCConfig `json:"rpc,omitempty"`

	// Observer is the configuration of the BT LE observer.
	Observer *BLEObserverConfig `json:"observer,omitempty"`
}

type BLERPCConfig struct {
	// Enable is true if rpc service is enabled, false otherwise.
	Enable *bool `json:"enable,omitempty"`
}

type BLEObserverConfig struct {
	// Enable is true if BT LE observer is enabled, false otherwise. Not applicable
	// for battery-operated devices.
	Enable *bool `json:"enable,omitempty"`
}

type BLEGetConfigRequest struct{}

func (r *BLEGetConfigRequest) Method() string {
	return "BLE.GetConfig"
}

func (r *BLEGetConfigRequest) NewTypedResponse() *BLEConfig {
	return &BLEConfig{}
}

func (r *BLEGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BLEGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*BLEConfig,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

type BLESetConfigRequest struct {
	Config BLERPCConfig `json:"config"`
}

func (r *BLESetConfigRequest) Method() string {
	return "BLE.SetConfig"
}

func (r *BLESetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *BLESetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BLESetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*SetConfigResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}
