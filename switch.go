package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

type SwitchGetConfigRequest struct {
	// ID of the switch component instance.
	ID int `json:"id"`
}

func (r *SwitchGetConfigRequest) Method() string {
	return "Switch.GetConfig"
}

func (r *SwitchGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*SwitchConfig,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

func (r *SwitchGetConfigRequest) NewTypedResponse() *SwitchConfig {
	return &SwitchConfig{}
}

func (r *SwitchGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type SwitchSetConfigRequest struct {
	// ID of the switch component instance.
	ID int `json:"id"`

	// Config that the method takes.
	Config SwitchConfig `json:"config"`
}

func (r *SwitchSetConfigRequest) Method() string {
	return "Switch.SetConfig"
}

func (r *SwitchSetConfigRequest) Do(
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

func (r *SwitchSetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *SwitchSetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type SwitchGetStatusRequest struct {
	// ID of the switch component instance.
	ID int `json:"id"`
}

func (r *SwitchGetStatusRequest) Method() string {
	return "Switch.GetStatus"
}

func (r *SwitchGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*SwitchStatus,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

func (r *SwitchGetStatusRequest) NewTypedResponse() *SwitchStatus {
	return &SwitchStatus{}
}

func (r *SwitchGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type SwitchSetRequest struct {
	// ID of the switch component instance.
	ID int `json:"id"`

	// On is true for switch on, false otherwise. Required
	On bool `json:"on"`

	// ToggleAfter is the number of seconds afterwhich the switch will flip-back.s
	ToggleAfter *float64 `json:"toggle_after,omitempty"`
}

func (r *SwitchSetRequest) Method() string {
	return "Switch.Set"
}

func (r *SwitchSetRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*SwitchActionResponse,
	*frame.Response,
	error,
) {
	resp := &SwitchActionResponse{}
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

func (r *SwitchSetRequest) NewTypedResponse() *SwitchActionResponse {
	return &SwitchActionResponse{}
}

func (r *SwitchSetRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type SwitchToggleRequest struct {
	// ID of the switch component instance.
	ID int `json:"id"`
}

func (r *SwitchToggleRequest) Method() string {
	return "Switch.Toggle"
}

func (r *SwitchToggleRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*SwitchActionResponse,
	*frame.Response,
	error,
) {
	resp := &SwitchActionResponse{}
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

func (r *SwitchToggleRequest) NewTypedResponse() *SwitchActionResponse {
	return &SwitchActionResponse{}
}

func (r *SwitchToggleRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type SwitchConfig struct {
	// ID of the switch component instance.
	ID int `json:"id"`

	// Name of the switch instance.
	Name *string `json:"name"`

	// InMode is the mode of the associated input. Range of values: momentary,
	// follow, flip, detached, cycle (if applicable), activate (if applicable)
	InMode *string `json:"in_mode,omitempty"`

	// InitialState is the output state to set on power_on. Range of values: off,
	// on, restore_last, match_input
	InitialState *string `json:"initial_state,omitempty"`

	// AutoOn is true if the "Automatic ON" function is enabled, false otherwise.
	AutoOn *bool `json:"auto_on,omitempty"`

	// AutoOnDelay is the number of seconds to pass until the component is
	// switched back.
	AutoOnDelay *float64 `json:"auto_on_delay,omitempty"`

	// AutoOff is true if the "Automatic OFF" function is enabled, false otherwise.
	AutoOff *bool `json:"auto_off,omitempty"`

	// AutoOffDelay is the number of seconds to pass until the component is switched back off.
	AutoOffDelay *float64 `json:"auto_off_delay,omitempty"`

	// AutorecoverVoltageErrors is true if switch output state should be restored
	// after over/undervoltage error is cleared, false otherwise (shown if applicable).
	AutorecoverVoltageErrors *bool `json:"autorecover_voltage_errors,omitempty"`

	// InputID is the ID of the Input component which controls the Switch.
	// Applicable only to Pro1 and Pro1PM devices. Valid values: 0, 1
	InputID *int `json:"input_id,omitempty"`

	// PowerLimit (in Watts) over which overpower condition occurs (shown if applicable).
	PowerLimit *float64 `json:"power_limit,omitempty"`

	// VoltageLimit (in Volts) over which overvoltage condition occurs (shown if applicable).
	VoltageLimit *float64 `json:"voltage_limit,omitempty"`

	// UndervoltageLimit (in Volts) over which overvoltage condition occurs (shown if applicable)
	UndervoltageLimit *float64 `json:"undervoltage_limit,omitempty"`

	// CurrentLimit (in Amperes) over which overcurrent condition occurs (shown if applicable)
	CurrentLimit *float64 `json:"current_limit,omitempty"`
}

type SwitchStatus struct {
	// ID of the switch component instance.
	ID int `json:"id"`

	// Source of the last command, for example: init, WS_in, http, ...
	Source *string `json:"source,omitempty"`

	// Output is true if the output channel is currently on, false otherwise.
	Output *bool `json:"output,omitempty"`

	// TimerStartedAt is the unix timestamp, start time of the timer (in UTC)
	// (shown if the timer is triggered)
	TimerStartedAt *float64 `json:"timer_started_at,omitempty"`

	// TimerDuration is the number of seconds for the timer (shown if the timer
	// is triggered).
	TimerDuration *float64 `json:"timer_duration,omitempty"`

	// APower is the last measured instantaneous active power (in Watts)
	// delivered to the attached load (shown if applicable).
	APower *float64 `json:"apower,omitempty"`

	// Voltage last measured in Volts (shown if applicable).
	Voltage *float64 `json:"voltage,omitempty"`

	// Current last measured in Amperes (shown if applicable).
	Current *float64 `json:"current,omitempty"`

	// PF is the last measured power factor (shown if applicable).
	PF *float64 `json:"pf,omitempty"`

	// Freq is the last measured network frequency in Hz (shown if applicable).
	Freq *float64 `json:"freq,omitempty"`

	// AEnergy contains information about the active energy counter (shown if
	// applicable)
	AEnergy *EnergyCounters `json:"aenergy,omitempty"`

	// RetAEnergy contains information about the returned active energy counter
	// (shown if applicable)
	RetAEnergy *EnergyCounters `json:"ret_aenergy,omitempty"`

	// Temperature describes the internal temperature of the relay.
	Temperature *Temperature `json:"temperature,omitempty"`

	// Errors lists error conditions occurred. May contain overtemp, overpower,
	// overvoltage, undervoltage, (shown if at least one error is present).
	Errors []string `json:"errors,omitempty"`
}

// EnergyCounters describes active energy counters.
type EnergyCounters struct {
	// Total energy consumed in Watt-hours.
	Total float64 `json:"total"`

	// ByMinute is the energy consumption by minute (in Milliwatt-hours) for
	// the last three minutes (the lower the index of the element in the array,
	// the closer to the current moment the minute)
	ByMinute []float64 `json:"by_minute"`

	// MinuteTS is the Unix timestamp of the first second of the last minute (in UTC)
	MinuteTS float64 `json:"minute_ts,omitempty"`
}

// Temperature describes a temperature measurement.
type Temperature struct {
	// C is the temperature in Celsius (null if temperature is out of the
	// measurement range)
	C *float64 `json:"tC,omitempty"`
	// F is the temperature in Fahrenheit (null if temperature is out of the
	// measurement range)
	F *float64 `json:"tF,omitempty"`
}

type SwitchActionResponse struct {
	// WasOn is true if the switch was on before the method was executed,
	// false otherwise.
	WasOn bool `json:"was_on"`
}
