package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

// LightGetConfigRequest contains parameters for the Light.GetConfig RPC request.
type LightGetConfigRequest struct {
	// ID of the light component instance.
	ID int `json:"id"`
}

func (r *LightGetConfigRequest) Method() string {
	return "Light.GetConfig"
}

func (r *LightGetConfigRequest) NewTypedResponse() *LightConfig {
	return &LightConfig{}
}

func (r *LightGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *LightGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*LightConfig,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// LightSetConfigRequest contains parameters for the Light.SetConfig RPC request.
type LightSetConfigRequest struct {
	// ID of the light component instance.
	ID int `json:"id"`

	// Config that the method takes.
	Config LightConfig `json:"config"`
}

func (r *LightSetConfigRequest) Method() string {
	return "Light.SetConfig"
}

func (r *LightSetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *LightSetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *LightSetConfigRequest) Do(
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

// LightGetStatusRequst contains parameters for the Light.GetStatus RPC request.
type LightGetStatusRequest struct {
	// ID of the light component instance.
	ID int `json:"id"`
}

func (r *LightGetStatusRequest) Method() string {
	return "Light.GetStatus"
}

func (r *LightGetStatusRequest) NewTypedResponse() *LightStatus {
	return &LightStatus{}
}

func (r *LightGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *LightGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*LightStatus,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// LightSetRequest is the parameters for the Light.Set RPC, which enables or disables a light.
type LightSetRequest struct {
	// ID of the light component instance.
	ID int `json:"id"`

	// On is true for light on, false otherwise. (optional). On or Brightness must be provided.
	On *bool `json:"on,omitempty"`

	// Brightness level (optional). On or Brightness must be provided.
	Brightness *float64 `json:"brightness,omitempty"`

	// TransitionDuration in seconds - time between change from current brightness level
	// to desired brightness level in request. (optional)
	TransitionDuration *float64 `json:"transition_duration,omitempty"`

	// ToggleAfter is the number of seconds afterwhich the light will flip-back. (optional)
	ToggleAfter *float64 `json:"toggle_after,omitempty"`
}

func (r *LightSetRequest) Method() string {
	return "Light.Set"
}

func (r *LightSetRequest) NewTypedResponse() *LightSetResponse {
	return &LightSetResponse{}
}

func (r *LightSetRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *LightSetRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*LightSetResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// LightSetResponse is the response body for the Light.Set RPC.
type LightSetResponse struct{}

// LightToggleRequest contains parameters for the Light.Toggle RPC request.
type LightToggleRequest struct {
	// ID of the light component instance.
	ID int `json:"id"`
}

func (r *LightToggleRequest) Method() string {
	return "Light.Toggle"
}

func (r *LightToggleRequest) NewTypedResponse() *LightToggleResponse {
	return &LightToggleResponse{}
}

func (r *LightToggleRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *LightToggleRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*LightToggleResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// LightToggleResponse is the body for the Light.Toggle RPC response.
type LightToggleResponse struct{}

// LightConfig provides configuration for light component instances.
type LightConfig struct {
	// ID of the light component instance.
	ID int `json:"id"`

	// Name of the light instance.
	Name *string `json:"name"`

	// InMode is the mode of the associated input. Range of values: follow, flip,
	// activate, detached, dim (if applicable), dual_dim (if applicable).
	InMode *string `json:"in_mode,omitempty"`

	// InitialState is the output state to set on power_on. Range of values: off,
	// on, restore_last, match_input
	InitialState *string `json:"initial_state,omitempty"`

	// AutoOn is true if the "Automatic ON" function is enabled, false otherwise.
	AutoOn *bool `json:"auto_on,omitempty"`

	// AutoOnDelay is the number of seconds to pass until the component is
	// lighted back.
	AutoOnDelay *float64 `json:"auto_on_delay,omitempty"`

	// AutoOff is true if the "Automatic OFF" function is enabled, false otherwise.
	AutoOff *bool `json:"auto_off,omitempty"`

	// AutoOffDelay is the number of seconds to pass until the component is lighted back off.
	AutoOffDelay *float64 `json:"auto_off_delay,omitempty"`

	// TransitionDuration (in seconds) - time to change from 0% to 100% of brightness (if
	// applicable).
	TransitionDuration *float64 `json:"transition_duration,omitempty"`

	// MinBrightnessOnToggle is the brightness level (in percent) applied when there is
	// a toggle and current brightness is lower than min_brightness_on_toggle.
	MinBrightnessOnToggle *float64 `json:"min_brightness_on_toggle,omitempty"`

	// NightMode configures the night mode feature.
	NightMode *LightNightModeConfig `json:"night_mode,omitempty"`

	// ButtonFadeRate controls how quickly the output level changes while a button is
	// held down for dimming (if applicable). Default value 3. Range [1,5] where 5 is
	// fastest, 1 is slowest.
	ButtonFadeRate *float64 `json:"button_fade_rate,omitempty"`

	// ButtonPresets provides configuration for button presets.
	ButtonPresets *LightButtonPresetsConfig `json:"button_presets,omitempty"`

	// RangeMap remaps output 0%-100% range to values in array (if applicable). First
	// value in array is min setting, second value is max setting. Array elements are
	// of type number. Float values are supported. Accepted range for values is from
	// 0% to 100%. Default values are [0, 100]. max must be greater than min.
	// nil may be specified to reset.
	RangeMap *[]float64 `json:"range_map,omitempty"`
}

// LightNightModeConfig configures the night mode feature.
type LightNightModeConfig struct {
	// Enable or disable night mode.
	Enable *bool `json:"enable,omitempty"`

	// Brightness level limit when night mode is active.
	Brightness *float64 `json:"brightness,omitempty"`

	// ActiveBetween is a slice containing 2 elements of type string, the first element
	// indicates the start of the period during which the night mode will be active,
	// the second indicates the end of that period. Both start and end are strings in
	// the format HH:MM, where HH and MM are hours and minutes with optinal leading zeros.
	ActiveBetween []string `json:"active_between,omitempty"`
}

// LightButtonPresetsConfig provides configuration for button presets.
type LightButtonPresetsConfig struct {
	// ButtonDoublePush configures button double push behavior. nil disables button_doublepush.
	ButtonDoublePush *LightButtonPresetsDoublePushConfig `json:"button_doublepush,omitempty"`
}

// LightButtonPresetsDoublePushConfig configures button double push behavior.
type LightButtonPresetsDoublePushConfig struct {
	// Brightness level (in percent) set on double click (if applicable), default: 100
	Brightness *float64 `json:"brightness,omitempty"`
}

// LightStatus describes the status of light component instances.
type LightStatus struct {
	// ID of the light component instance.
	ID int `json:"id"`

	// Source of the last command, for example: init, WS_in, http, ...
	Source *string `json:"source,omitempty"`

	// Output is true if the output channel is currently on, false otherwise.
	Output *bool `json:"output,omitempty"`

	// Brightness level (in percent)
	Brightness *float64 `json:"brightness,omitempty"`

	// TimerStartedAt is the unix timestamp, start time of the timer (in UTC)
	// (shown if the timer is triggered)
	TimerStartedAt *float64 `json:"timer_started_at,omitempty"`

	// TimerDuration is the number of seconds for the timer (shown if the timer
	// is triggered).
	TimerDuration *float64 `json:"timer_duration,omitempty"`

	// Transition provides information about the transition (shown if transition is triggered).
	Transition *LightTransitionStatus `json:"transition,omitempty"`

	// Temperature describes the internal temperature of the relay.
	Temperature *Temperature `json:"temperature,omitempty"`
}

// LightTransitionStatus provides information about the transition (shown if transition
// is triggered).
type LightTransitionStatus struct {
	// Target describes the desired result of the transition.
	Target LightTransitionTargetStatus `json:"target,omitempty"`

	// StartedAt is the unix timestamp start time of the transition (in UTC).
	StartedAt *float64 `json:"started_at,omitempty"`

	// Duration of the transition in seconds.
	Duration *float64 `json:"duration,omitempty"`
}

// LightTransitionTargetStatus describes the desired result of the transition.
type LightTransitionTargetStatus struct {
	// Output is true if the output channel becomes on, false otherwise
	Output bool `json:"output"`

	// Brightness level (in percent).
	Brightness *float64 `json:"brightness,omitempty"`
}
