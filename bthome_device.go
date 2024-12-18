package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

type BTHomeDeviceConfig struct {
	// ID of the component instance.
	ID int `json:"id"`

	// Name of the component instance.
	Name *string `json:"name"`

	// MAC address of the physical device
	Addr string `json:"addr"`

	// AES encryption key as hexadecimal string for encrypted devices
	Key *string `json:"key"`

	// Meta contains meta data for the component.
	Meta BTHomeDeviceConfigMeta `json:"meta"`
}

// Object for storing meta data
type BTHomeDeviceConfigMeta struct {
	// UI contains setting for how the component will be rendered in the UI.
	UI BTHomeDeviceConfigMetaUI `json:"ui"`
}

// BTHomeDeviceConfigMetaUI contains setting for how the component will be rendered in the UI.
type BTHomeDeviceConfigMetaUI struct {
	// Icon allows setting custom icon for the component's card by providing an external hosted image via link.
	Icon *string `json:"icon"`
}

type BTHomeDeviceStatus struct {
	// ID of the component instance.
	ID int `json:"id"`

	// RSSI is the Strength of the signal in dBms from the latest packet
	RSSI *float64 `json:"rssi,omitempty"`

	// Battery is the battery level in percentage
	Battery *float64 `json:"battery,omitempty"`

	// PacketID is the ID of the latest received packet
	PacketID *int `json:"packet_id,omitempty"`

	// LastUpdateTS is the timestamp of the received packet.
	LastUpdateTS int `json:"last_update_ts"`

	// Errors describes component error conditions. May contain key_missing_or_bad, decrypt_failed,
	// parse_failed and unencrypted_data.
	Errors []string `json:"errors,omitempty"`
}

// BTHomeDeviceGetConfigRequest contains parameters for the BTHomeDevice.GetConfig RPC request.
type BTHomeDeviceGetConfigRequest struct {
	// ID of the component instance.
	ID int `json:"id"`
}

func (r *BTHomeDeviceGetConfigRequest) Method() string {
	return "BTHomeDevice.GetConfig"
}

func (r *BTHomeDeviceGetConfigRequest) NewTypedResponse() *BTHomeDeviceConfig {
	return &BTHomeDeviceConfig{}
}

func (r *BTHomeDeviceGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BTHomeDeviceGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*BTHomeDeviceConfig,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

// BTHomeDeviceSetConfigRequest contains parameters for the BTHomeDevice.SetConfig RPC request.
type BTHomeDeviceSetConfigRequest struct {
	// ID of the component instance.
	ID int `json:"id"`

	// Config that the method takes.
	Config BTHomeDeviceConfig `json:"config"`
}

func (r *BTHomeDeviceSetConfigRequest) Method() string {
	return "BTHomeDevice.SetConfig"
}

func (r *BTHomeDeviceSetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *BTHomeDeviceSetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BTHomeDeviceSetConfigRequest) Do(
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

// BTHomeDeviceGetStatusRequst contains parameters for the BTHomeDevice.GetStatus RPC request.
type BTHomeDeviceGetStatusRequest struct {
	// ID of the component instance.
	ID int `json:"id"`
}

func (r *BTHomeDeviceGetStatusRequest) Method() string {
	return "BTHomeDevice.GetStatus"
}

func (r *BTHomeDeviceGetStatusRequest) NewTypedResponse() *BTHomeDeviceStatus {
	return &BTHomeDeviceStatus{}
}

func (r *BTHomeDeviceGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BTHomeDeviceGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*BTHomeDeviceStatus,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

// BTHomeDeviceGetKnownObjectsRequst contains parameters for the BTHomeDevice.GetKnownObjects RPC request.
type BTHomeDeviceGetKnownObjectsRequest struct {
	// ID of the component instance.
	ID int `json:"id"`
}

func (r *BTHomeDeviceGetKnownObjectsRequest) Method() string {
	return "BTHomeDevice.GetKnownObjects"
}

func (r *BTHomeDeviceGetKnownObjectsRequest) NewTypedResponse() *BTHomeDeviceGetKnownObjectsResponse {
	return &BTHomeDeviceGetKnownObjectsResponse{}
}

func (r *BTHomeDeviceGetKnownObjectsRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BTHomeDeviceGetKnownObjectsRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*BTHomeDeviceGetKnownObjectsResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type BTHomeDeviceGetKnownObjectsResponse struct {
	// ID of the component instance.
	ID int `json:"id"`

	// Objects is a list of known objects
	Objects []BTHomeDeviceKnownObject `json:"objects"`
}

type BTHomeDeviceKnownObject struct {
	// ObjID is the BTHome object id in decimal
	ObjID int `json:"obj_id"`

	// IDX is the BTHome object index in decimal
	IDX int `json:"idx"`

	// Component key if the sensor is managed, otherwise nil.
	Component *string `json:"component,omitempty"`
}
