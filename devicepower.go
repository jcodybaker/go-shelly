package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

type DevicePowerStatus struct {
	// ID of the devicepower component instance.
	ID int `json:"id"`

	// Battery is information about the battery charge.
	Battery *DevicePowerBatteryStatus `json:"battery,omitempty"`

	// External is about the external power source (only available if external power source is supported).
	External *DevicePowerExternalStatus `json:"external,omitempty"`

	// Errors is a list of error events related to device power.
	Errors []string `json:"errors,omitempty"`
}

// DevicePowerBatteryStatus is information about the battery charge.
type DevicePowerBatteryStatus struct {
	// V is battery voltage in Volts (null if valid value could not be obtained).
	V *float64 `json:"V,omitempty"`

	// Percent is the battery charge level in % (null if valid value could not be obtained).
	Percent *float64 `json:"percent,omitempty"`
}

// DevicePowerExternalStatus is about the external power source (only available if external power source is supported).
type DevicePowerExternalStatus struct {
	// Present is true if external power source is connected, false otherwise.
	Present bool `json:"present"`
}

type DevicePowerGetStatusRequest struct {
	// ID of the DevicePower component instance.
	ID int `json:"id"`
}

func (r *DevicePowerGetStatusRequest) Method() string {
	return "DevicePower.GetStatus"
}

func (r *DevicePowerGetStatusRequest) NewTypedResponse() *DevicePowerStatus {
	return &DevicePowerStatus{}
}

func (r *DevicePowerGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *DevicePowerGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*DevicePowerStatus,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}
