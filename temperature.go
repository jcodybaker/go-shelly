package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

// TemperatureGetConfigRequest contains parameters for the Temperature.GetConfig RPC request.
type TemperatureGetConfigRequest struct {
	// ID of the temperature component instance.
	ID int `json:"id"`
}

func (r *TemperatureGetConfigRequest) Method() string {
	return "Temperature.GetConfig"
}

func (r *TemperatureGetConfigRequest) NewTypedResponse() *TemperatureConfig {
	return &TemperatureConfig{}
}

func (r *TemperatureGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *TemperatureGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*TemperatureConfig,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

// TemperatureSetConfigRequest contains parameters for the Temperature.SetConfig RPC request.
type TemperatureSetConfigRequest struct {
	// ID of the temperature component instance.
	ID int `json:"id"`

	// Config that the method takes.
	Config TemperatureConfig `json:"config"`
}

func (r *TemperatureSetConfigRequest) Method() string {
	return "Temperature.SetConfig"
}

func (r *TemperatureSetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *TemperatureSetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *TemperatureSetConfigRequest) Do(
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

// TemperatureGetStatusRequst contains parameters for the Temperature.GetStatus RPC request.
type TemperatureGetStatusRequest struct {
	// ID of the temperature component instance.
	ID int `json:"id"`
}

func (r *TemperatureGetStatusRequest) Method() string {
	return "Temperature.GetStatus"
}

func (r *TemperatureGetStatusRequest) NewTypedResponse() *TemperatureStatus {
	return &TemperatureStatus{}
}

func (r *TemperatureGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *TemperatureGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*TemperatureStatus,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

// TemperatureConfig provides configuration for temperature component instances.
type TemperatureConfig struct {
	// ID of the temperature component instance.
	ID int `json:"id"`

	// Name of the temperature instance.
	Name *string `json:"name"`

	// ReportTHR is the temperature report threshold in Celsius. Accepted range is device-specific,
	// default [0.5..5.0]C unless specified otherwise.
	ReportTHR *float64 `json:"report_thr_C,omitempty"`

	// OffsetC is the offset in Celsius to be applied to the measured temperature. Accepted range is
	// device-specific, default [-50.0 .. 50.0] unless specified otherwise.
	OffsetC *float64 `json:"offset_C,omitempty"`
}

// TemperatureStatus describes the status of temperature component instances.
type TemperatureStatus struct {
	// ID of the temperature component instance.
	ID int `json:"id"`

	// TC is the temperature in Celsius (null if valid value could not be obtained)
	TC *float64 `json:"tC,omitempty"`

	// TF is the temperature in Fahrenheit  (null if valid value could not be obtained)
	TF *float64 `json:"tF,omitempty"`

	// Errors is a list of error shown only if at least one error is present. May contain
	// out_of_range, read when there is problem reading sensor.
	Errors []string `json:"errors,omitempty"`
}
