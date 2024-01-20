package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

type CoverGetConfigRequest struct {
	// ID of the cover component instance.
	ID int `json:"id"`
}

func (r *CoverGetConfigRequest) Method() string {
	return "Cover.GetConfig"
}

func (r *CoverGetConfigRequest) NewTypedResponse() *CoverConfig {
	return &CoverConfig{}
}

func (r *CoverGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CoverGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*CoverConfig,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

type CoverSetConfigRequest struct {
	// ID of the cover component instance.
	ID int `json:"id"`

	// Config that the method takes.
	Config CoverConfig `json:"config"`
}

func (r *CoverSetConfigRequest) Method() string {
	return "Cover.SetConfig"
}

func (r *CoverSetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *CoverSetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CoverSetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*CoverConfig,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

type CoverConfig struct {
	// ID of the cover component instance.
	ID int `json:"id"`

	// Name of the cover instance.
	Name *string `json:"name"`

	// InMode is the mode of the associated input. One of single, dual or detached,
	// only present if there is at least one input associated with the Cover instance.
	//
	// `single` - Cover operation in both open and close directions is controlled via
	//    a single input. In this mode, only input_0 is used to open/close/stop the Cover.
	//    It doesn't matter if input_0 has in_type=switch or in_type=button, the behavior is
	//    the same: each switch toggle or button press cycles between open/stop/close/stop/...
	//    In single mode, input_1 is free to be used as a safety switch (e.g. end-of-motion
	//    limit switch, emergency-stop, etc.), see below.
	// `dual` - Cover operation is controlled via two inputs, one for open and one for close.
	//    In this mode, input_0 is used to open the Cover, input_1 is used to close the
	//    Cover.The exact behavior depends on the in_type of the inputs: if in_type = switch:
	//    toggle the switch to ON to move in the associated direction; toggle the switch to
	//    OFF to stop, if in_type = button: press the button to move in the associated
	//    direction; press the button again to stop.
	// `detached` - Cover operation via the input/inputs is prohibited.
	InMode *string `json:"in_mode,omitempty"`

	// InitialState defines Cover target state on power-on, one of open (Cover will fully
	// open), closed (Cover will fully close) or stopped (Cover will not change its position).
	InitialState *string `json:"initial_state,omitempty"`

	// PowerLimit (in Watts) over which overpower condition occurs (shown if applicable).
	PowerLimit *float64 `json:"power_limit,omitempty"`

	// VoltageLimit (in Volts) over which overvoltage condition occurs (shown if applicable).
	VoltageLimit *float64 `json:"voltage_limit,omitempty"`

	// UndervoltageLimit (in Volts) over which overvoltage condition occurs (shown if applicable)
	UndervoltageLimit *float64 `json:"undervoltage_limit,omitempty"`

	// CurrentLimit (in Amperes) over which overcurrent condition occurs (shown if applicable)
	CurrentLimit *float64 `json:"current_limit,omitempty"`

	// Motor is the configuration of the Cover motor. The exact contents depend on the type of
	// motor used. The descriptions below are valid when an AC motor is use.
	Motor *CoverMotorConfig `json:"motor,omitempty"`

	// MaxTimeOpen is the timeout after which Cover will stop moving in open direction.
	MaxTimeOpen *float64 `json:"max_time_open,omitempty"`

	// MaxTimeClose is the timeout after which Cover will stop moving in a close direction.
	MaxTimeClose *float64 `json:"max_time_close,omitempty"`

	// SwapInputs is only present if there are two inputs associated with the Cover instance,
	// defines whether the functions of the two inputs are swapped. The effect of swap_inputs
	// is observable only when in_mode != detached.
	//
	// false - When swap_inputs is false: If in_mode = dual: input_0 is used to open, input_1
	//   is used to close. If in_mode = single: input_0 is used to open/close/stop, input_1 is
	//   used as a safety switch or is not used at all.
	// true - When swap_inputs is true: If in_mode = dual: input_0 is used to close, input_1
	//   is used to open. If in_mode = single: input_0 is used as a safety switch or is not
	//   used at all, input_1 is used to open/close/stop.
	SwapInputs *bool `json:"swap_inputs,omitempty"`

	// InvertDirections defines the motor rotation for open and close directions (changing
	// this parameter requires a reboot).
	//
	// false - On open motor rotates clockwise, on close motor rotates counter-clockwise.
	// true - On open motor rotates counter-clockwise, on close motor rotates clockwise.
	InvertDirections *bool `json:"invert_directions,omitempty"`

	// ObstructionDetection defines the behavior of the obstruction detection safety feature.
	ObstructionDetection *CoverObstructionDetectionConfig `json:"obstruction_detection,omitempty"`

	// SafetySwitch defines the behavior of the safety switch feature, only present
	// if there are two inputs associated with the Cover instance. The safety_switch feature
	// will only work when in_mode=single.
	SafetySwitch *CoverSafetySwitchConfig `json:"safety_switch,omitempty"`
}

// CoverMotorConfig provides configuration for the Cover motor.
type CoverMotorConfig struct {
	// IdlePowerThr is the threshold in watts, below which the motor is considered stopped.
	IdlePowerThr *float64 `json:"idle_power_thr,omitempty"`

	// IdleConfirmPeriod is the minimum period of time (seconds) in idle state before the
	// state is confirmed.
	IdleConfirmPeriod *float64 `json:"idle_confirm_period,omitempty"`
}

// CoverObstructionDetectionConfig defines the behavior of the obstruction detection
// safety feature.
type CoverObstructionDetectionConfig struct {
	// Enable is true when obstruction detection is enabled, false otherwise.
	Enable *bool `json:"enable,omitempty"`

	// Direction of motion for which the safety switch should be monitored, one
	// of open, close, both.
	Direction *string `json:"direction,omitempty"`

	// Action (for recovery) that should be performed if the safety switch is engaged
	// while moving in a monitored direction, one of:
	//
	// stop - Immediately stop Cover.
	// reverse - Immediately stop Cover, then move in the opposite direction until a
	//   fully open or fully closed position is reached. If Cover encounters a new
	//   obstruction while reversing from a previous one, it will unconditionally stop.
	Action *string `json:"action,omitempty"`

	// PowerThr is the threshhold in watts, which should be interpreted as objects
	// obstructing Cover movement. This property is editable at any time, but note that
	// during the cover calibration procedure (Cover.Calibrate), power_thr will be
	// automatically set to the peak power consumption + 15%, overwriting the current
	// value. The automatic setup of power_thr during calibration will only start tracking
	// power values when the holdoff time (see below) has elapsed.
	PowerThr *float64 `json:"power_thr,omitempty"`

	// Seconds, time to wait after Cover starts moving before obstruction detection is
	// activated (to avoid false detections because of the initial power consumption spike).
	HoldOff *float64 `json:"holdoff,omitempty"`
}

// CoverSafetySwitchConfig defines the behavior of the safety switch feature, only present
// if there are two inputs associated with the Cover instance. The safety_switch feature
// will only work when in_mode=single.
type CoverSafetySwitchConfig struct {
	// Enable is true when saftey switch is enabled, false otherwise.
	Enable *bool `json:"enable,omitempty"`

	// Direction of motion for which the safety switch should be monitored, one
	// of open, close, both.
	Direction *string `json:"direction,omitempty"`

	// Action (for recovery) that should be performed if the safety switch is engaged
	// while moving in a monitored direction, one of:
	//
	// stop - Immediately stop Cover, then wait for a command to move in an allowed
	//   direction (see below).
	// reverse - Immediately stop Cover, then move in the opposite direction until a fully
	//   open or fully closed position is reached. action = reverse requires that
	//   allowed_move = reverse.
	// pause - Immediately stop Cover, then either: wait for a command to move in an
	//   allowed direction (see below) or automatically continue movement in the same
	//   direction (i.e. the one that was interrupted) when the safety switch is disengaged.
	Action *string `json:"action,omitempty"`

	// AllowedMove direction when the safety switch is engaged while moving in a monitored
	// direction:
	//
	// nil - null means Cover can't be moved in either open nor closed directions while
	//   the safety switch is engaged.
	// reverse - The only other option is reverse, which means Cover can only be moved in
	//   the direction opposite to the one that was interrupted (for example, if the safety
	//   switch was hit while opening, Cover can only be commanded to close if the switch is
	//   not disengaged)
	AllowedMove *string `json:"allowed_move,omitempty"`
}

type CoverGetStatusRequest struct {
	// ID of the cover component instance.
	ID int `json:"id"`
}

func (r *CoverGetStatusRequest) Method() string {
	return "Cover.GetStatus"
}

func (r *CoverGetStatusRequest) NewTypedResponse() *CoverStatus {
	return &CoverStatus{}
}

func (r *CoverGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CoverGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*CoverStatus,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// CoverStatus describes the current state of the Cover.
type CoverStatus struct {
	// ID of the cover component instance.
	ID int `json:"id"`

	// Source of the last command, for example: init, WS_in, http, ...
	Source *string `json:"source,omitempty"`

	// State describes the current state of the cover device. One of open (Cover is
	// fully open), closed (Cover is fully closed), opening (Cover is actively opening),
	// closing (Cover is actively closing), stopped (Cover is not moving, and is neither
	// fully open nor fully closed, or the open/close state is unknown), calibrating
	// (Cover is performing a calibration procedure).
	State *string `json:"state,omitempty"`

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

	// CurrentPos is the current position current position in percent from 0 (fully
	// closed) to 100 (fully open); null if the position is unknown. Only present if
	// Cover is calibrated.
	CurrentPos *float64 `json:"current_pos,omitempty"`

	// TargetPos is only present if Cover is calibrated and is actively moving to a
	// requested position in either open or closed directions. Represents the target
	// position in percent from 0 (fully closed) to 100 (fully open); null if target
	// position has been reached or the movement was canceled.
	TargetPos *float64 `json:"target_pos,omitempty"`

	// MoveTimeout is the timeout in seconds until the cover stops regardless of completion.
	// Only present if Cover is actively moving in either open or closed directions.
	MoveTimeout *float64 `json:"move_timeout,omitempty"`

	// MoveStartedAt represents the time at which the movement has begun. Only present if
	// Cover is actively moving in either open or closed directions.
	MoveStartedAt *float64 `json:"move_started_at,omitempty"`

	// PosControl is false if Cover is not calibrated and only discrete open/close is
	// possible; true if Cover is calibrated and can be commanded to go to arbitrary
	// positions between fully open and fully closed.
	PosControl *bool `json:"pos_control,omitempty"`

	// LastDirection is the direction of the last movement: open/close or null when unknown.
	LastDirection *string `json:"last_direction,omitempty"`

	// Temperature describes the internal temperature of the cover instance. Only present if
	// a temperature monitor is associated with the Cover instance
	Temperature *Temperature `json:"temperature,omitempty"`

	// Errors lists error conditions occurred. May contain overtemp, overpower,
	// overvoltage, undervoltage, (shown if at least one error is present).
	Errors []string `json:"errors,omitempty"`
}

// CoverCalibrateRequest causes the device to enter calibration mode. See:
// - https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Cover#covercalibrate
// - https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Cover#calibration-kb
type CoverCalibrateRequest struct {
	// ID of the cover component instance.
	ID int `json:"id"`
}

func (r *CoverCalibrateRequest) Method() string {
	return "Cover.Calibrate"
}

func (r *CoverCalibrateRequest) NewTypedResponse() *CoverCalibrateRespose {
	return &CoverCalibrateRespose{}
}

func (r *CoverCalibrateRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CoverCalibrateRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*CoverCalibrateRespose,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// CoverCalibrateRespose is the RPC response for Cover.Calibrate.
type CoverCalibrateRespose struct{}

// CoverOpenRequest causes the device to open the cover instance.
type CoverOpenRequest struct {
	// ID of the cover component instance.
	ID int `json:"id"`

	// Duration (seconds) if provided, Cover will move in the open direction for the
	// specified time. duration must be in the range [0.1..maxtime_open].
	// If duration is not provided, Cover will fully open, unless it times out because
	// of maxtime_open first.
	Duration *float64 `json:"duration,omitempty"`
}

func (r *CoverOpenRequest) Method() string {
	return "Cover.Open"
}

func (r *CoverOpenRequest) NewTypedResponse() *CoverOpenResponse {
	return &CoverOpenResponse{}
}

func (r *CoverOpenRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CoverOpenRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*CoverOpenResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// CoverOpenResponse is the RPC response for Cover.Open.
type CoverOpenResponse struct{}

// CoverCloseRequest causes the device to close the cover instance.
type CoverCloseRequest struct {
	// ID of the cover component instance.
	ID int `json:"id"`

	// Duration (seconds) if provided, Cover will move in the close direction for the
	// specified time. duration must be in the range [0.1..maxtime_open].
	// If duration is not provided, Cover will fully close, unless it times out because
	// of maxtime_close first.
	Duration *float64 `json:"duration,omitempty"`
}

func (r *CoverCloseRequest) Method() string {
	return "Cover.Close"
}

func (r *CoverCloseRequest) NewTypedResponse() *CoverCloseResponse {
	return &CoverCloseResponse{}
}

func (r *CoverCloseRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CoverCloseRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*CoverCloseResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// CoverCloseResponse is the RPC response for Cover.Close.
type CoverCloseResponse struct{}

// CoverStopRequest causes the device to stop in progress actions for the cover instance.
type CoverStopRequest struct {
	// ID of the cover component instance.
	ID int `json:"id"`
}

func (r *CoverStopRequest) Method() string {
	return "Cover.Stop"
}

func (r *CoverStopRequest) NewTypedResponse() *CoverStopResponse {
	return &CoverStopResponse{}
}

func (r *CoverStopRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CoverStopRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*CoverStopResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// CoverStopResponse is the RPC response for Cover.Stop.
type CoverStopResponse struct{}

// CoverGoToPositionRequest causes the device to travel to the specified position.
type CoverGoToPositionRequest struct {
	// ID of the cover component instance.
	ID int `json:"id"`

	// Pos represents target position in %, allowed range [0..100].
	// Pos is mutually exclusive with Rel. Rel or Pos is required, but both may not be set.
	Pos *float64 `json:"pos,omitempty"`

	// Rel represents a relative move in %, allowed range [-100..100] Cover will move
	// to a target_position = current_position + rel. If the value of rel is so big that
	// it results in overshoot (i.e. target_position is beyond fully open / fully closed),
	// target_position will be silently capped to fully open / fully closed.
	// Rel is mutually exclusive with Pos. Rel or Pos is required, but both may not be set.
	Rel *float64 `json:"rel,omitempty"`
}

func (r *CoverGoToPositionRequest) Method() string {
	return "Cover.GoToPosition"
}

func (r *CoverGoToPositionRequest) NewTypedResponse() *CoverGoToPositionResponse {
	return &CoverGoToPositionResponse{}
}

func (r *CoverGoToPositionRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CoverGoToPositionRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*CoverGoToPositionResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// CoverGoToPositionResponse is the RPC response for Cover.GoToPosition.
type CoverGoToPositionResponse struct{}

// CoverResetCountersRequest resets counters for the cover.
type CoverResetCountersRequest struct {
	// ID of the cover component instance.
	ID int `json:"id"`

	// Type describes which counters should be reset.
	Type []string `json:"type,omitempty"`
}

func (r *CoverResetCountersRequest) Method() string {
	return "Cover.ResetCounters"
}

func (r *CoverResetCountersRequest) NewTypedResponse() *CoverResetCountersResponse {
	return &CoverResetCountersResponse{}
}

func (r *CoverResetCountersRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *CoverResetCountersRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*CoverResetCountersResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// CoverResetCountersResponse is the RPC response for Cover.ResetCounters.
type CoverResetCountersResponse struct {
	// AEnergy contains information about the active energy counter prior to reset.
	AEnergy *EnergyCounters `json:"aenergy,omitempty"`
}
