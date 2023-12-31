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

func (r *BLEGetStatusRequest) NewResponse() *BLEStatus {
	return &BLEStatus{}
}

func (r *BLEGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*BLEStatus,
	*frame.Response,
	error,
) {
	resp := r.NewResponse()
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

func (r *BLEGetConfigRequest) NewResponse() *BLEConfig {
	return &BLEConfig{}
}

func (r *BLEGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*BLEConfig,
	*frame.Response,
	error,
) {
	resp := r.NewResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

type BLESetConfigRequest struct{}

func (r *BLESetConfigRequest) Method() string {
	return "BLE.SetConfig"
}

func (r *BLESetConfigRequest) NewResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *BLESetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*SetConfigResponse,
	*frame.Response,
	error,
) {
	resp := r.NewResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}
