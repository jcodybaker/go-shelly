package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

// HumidityGetConfigRequest contains parameters for the Humidity.GetConfig RPC request.
type HumidityGetConfigRequest struct {
	// ID of the humidity component instance.
	ID int `json:"id"`
}

func (r *HumidityGetConfigRequest) Method() string {
	return "Humidity.GetConfig"
}

func (r *HumidityGetConfigRequest) NewTypedResponse() *HumidityConfig {
	return &HumidityConfig{}
}

func (r *HumidityGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *HumidityGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*HumidityConfig,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

// HumiditySetConfigRequest contains parameters for the Humidity.SetConfig RPC request.
type HumiditySetConfigRequest struct {
	// ID of the humidity component instance.
	ID int `json:"id"`

	// Config that the method takes.
	Config HumidityConfig `json:"config"`
}

func (r *HumiditySetConfigRequest) Method() string {
	return "Humidity.SetConfig"
}

func (r *HumiditySetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *HumiditySetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *HumiditySetConfigRequest) Do(
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

// HumidityGetStatusRequst contains parameters for the Humidity.GetStatus RPC request.
type HumidityGetStatusRequest struct {
	// ID of the humidity component instance.
	ID int `json:"id"`
}

func (r *HumidityGetStatusRequest) Method() string {
	return "Humidity.GetStatus"
}

func (r *HumidityGetStatusRequest) NewTypedResponse() *HumidityStatus {
	return &HumidityStatus{}
}

func (r *HumidityGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *HumidityGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*HumidityStatus,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

// HumidityConfig provides configuration for humidity component instances.
type HumidityConfig struct {
	// ID of the humidity component instance.
	ID int `json:"id"`

	// Name of the humidity instance.
	Name *string `json:"name"`

	// ReportTHR is the humidity report threshold in %. Accepted range is device-specific,
	// default [1.0..20.0]% unless specified otherwise.
	ReportTHR *float64 `json:"report_thr,omitempty"`

	// Offset in %. Value is applied to measured humidity. Accepted range is device-specific, default [-50.0..50.0]% unless specified otherwise
	Offset *float64 `json:"offset,omitempty"`
}

// HumidityStatus describes the status of humidity component instances.
type HumidityStatus struct {
	// ID of the humidity component instance.
	ID int `json:"id"`

	// RH is the relative humidity in % (null if valid value could not be obtained)
	RH *float64 `json:"rh,omitempty"`

	// Errors is a list of error events related to humidity.
	Errors []string `json:"errors,omitempty"`
}
