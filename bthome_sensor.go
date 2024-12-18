package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

type BTHomeSensorConfig struct {
	// ID of the component instance.
	ID int `json:"id"`

	// Name of the component instance.
	Name *string `json:"name"`

	// ObjID is the BTHome object id in decimal
	ObjID int `json:"obj_id"`

	// IDX is the BTHome object index in decimal
	IDX int `json:"idx"`

	// MAC address of the physical device
	Addr string `json:"addr"`

	// Meta contains meta data for the component.
	Meta BTHomeSensorConfigMeta `json:"meta"`
}

// Object for storing meta data
type BTHomeSensorConfigMeta struct {
	// UI contains setting for how the component will be rendered in the UI.
	UI BTHomeSensorConfigMetaUI `json:"ui"`
}

// BTHomeSensorConfigMetaUI contains setting for how the component will be rendered in the UI.
type BTHomeSensorConfigMetaUI struct {
	// Icon allows setting custom icon for the component's card by providing an external hosted image via link.
	Icon *string `json:"icon"`
}

// BTHomeSensorGetConfigRequest contains parameters for the BTHomeSensor.GetConfig RPC request.
type BTHomeSensorGetConfigRequest struct {
	// ID of the BTHomeSensor component instance.
	ID int `json:"id"`
}

func (r *BTHomeSensorGetConfigRequest) Method() string {
	return "BTHomeSensor.GetConfig"
}

func (r *BTHomeSensorGetConfigRequest) NewTypedResponse() *BTHomeSensorConfig {
	return &BTHomeSensorConfig{}
}

func (r *BTHomeSensorGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BTHomeSensorGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*BTHomeSensorConfig,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

// BTHomeSensorSetConfigRequest contains parameters for the BTHomeSensor.SetConfig RPC request.
type BTHomeSensorSetConfigRequest struct {
	// ID of the BTHomeSensor component instance.
	ID int `json:"id"`

	// Config that the method takes.
	Config BTHomeSensorConfig `json:"config"`
}

func (r *BTHomeSensorSetConfigRequest) Method() string {
	return "BTHomeSensor.SetConfig"
}

func (r *BTHomeSensorSetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *BTHomeSensorSetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BTHomeSensorSetConfigRequest) Do(
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

// BTHomeSensorGetStatusRequst contains parameters for the BTHomeSensor.GetStatus RPC request.
type BTHomeSensorGetStatusRequest struct {
	// ID of the BTHomeSensor component instance.
	ID int `json:"id"`
}

func (r *BTHomeSensorGetStatusRequest) Method() string {
	return "BTHomeSensor.GetStatus"
}

func (r *BTHomeSensorGetStatusRequest) NewTypedResponse() *BTHomeSensorStatus {
	return &BTHomeSensorStatus{}
}

func (r *BTHomeSensorGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *BTHomeSensorGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*BTHomeSensorStatus,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

// BTHomeSensorStatus describes the status of BTHomeSensor component instances.
type BTHomeSensorStatus struct {
	// ID of the BTHomeSensor component instance.
	ID int `json:"id"`

	// Value of the sensor (latest).
	Value interface{} `json:"value,omitempty"`

	// LastUpdateTS is the timestamp of the last received value.
	LastUpdateTS *float64 `json:"last_update_ts,omitempty"`
}

func (r *BTHomeSensorStatus) GetIntValue() (int, bool) {
	v, ok := r.Value.(int)
	return v, ok
}

func (r *BTHomeSensorStatus) GetFloatValue() (float64, bool) {
	v, ok := r.Value.(float64)
	return v, ok
}

func (r *BTHomeSensorStatus) GetStringValue() (string, bool) {
	v, ok := r.Value.(string)
	return v, ok
}