package shelly

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

type InputGetStatusRequest struct {
	// ID of the switch component instance.
	ID int `json:"id"`
}

func (r *InputGetStatusRequest) Method() string {
	return "Input.GetStatus"
}

func (r *InputGetStatusRequest) NewResponse() *InputStatus {
	return &InputStatus{}
}

func (r *InputGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*InputStatus,
	*frame.Response,
	error,
) {
	resp := r.NewResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

type InputStatus struct {
	// ID of the input component instance.
	ID int `json:"id"`

	// State of the input (null if the input instance is stateless, i.e. for type button)
	// (only for type switch, button).
	State *bool `json:"state,omitempty"`

	// Percent is the analog value in percent (null if the valid value could not be obtained)
	// (only for type "analog").
	Percent *float64 `json:"percent,omitempty"`

	// XPercent is percent transformed with config.xpercent.expr. Present only when both
	// config.xpercent.expr and config.xpercent.unit are set to non-empty values. null if
	// config.xpercent.expr can not be evaluated.
	// (only for type "analog").
	XPercent *float64 `json:"xpercent,omitempty"`

	// Errors is shown only if at least one error is present. May contain out_of_range, read.
	Errors []string `json:"errors,omitempty"`
}

type InputGetConfigRequest struct {
	// ID of the input component instance.
	ID int `json:"id"`
}

func (r *InputGetConfigRequest) Method() string {
	return "Input.GetConfig"
}

func (r *InputGetConfigRequest) NewResponse() *InputConfig {
	return &InputConfig{}
}

func (r *InputGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*InputConfig,
	*frame.Response,
	error,
) {
	resp := r.NewResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

type InputConfig struct {
	// ID of the switch component instance.
	ID int `json:"id"`

	// Name of the switch instance.
	Name *string `json:"name"`

	// Type of associated input. Range of values switch, button, analog (only if applicable).
	Type *string `json:"type,omitempty"`

	// Enable flag. When disabled, the input instance doesn't emit any events and reports
	// status properties as null. Applies for all input types.
	Enable *bool `json:"enable,omitempty"`

	// Invert is true if the logical state of the associated input is inverted, false otherwise.
	// For the change to be applied, the physical switch has to be toggled once after invert
	// is set. For type analog inverts percent range - 100% becomes 0% and 0% becomes 100%
	Invert *bool `json:"invert,omitempty"`

	// FactoryReset is true if input-triggered factory reset option is enabled, false otherwise
	// (shown if applicable). (only for type switch, button).
	FactorReset *bool `json:"factory_reset,omitempty"`

	// ReportThr is the analog input report threshold in percent. The accepted range is
	// device-specific, default [1.0..50.0]% unless specified otherwise.
	ReportThr *float64 `json:"report_thr,omitempty"`

	// RangeMap remaps 0%-100% range to values in array. The first value in the array is the
	// min setting, and the second value is the max setting. Array elements are of type number.
	// Float values are supported. The accepted range for values is from 0% to 100%. Default
	// values are [0, 100]. max must be greater than min. Equality is supported.
	RangeMap []float64 `json:"range_map,omitempty"`

	// XPercent is value transformation config for status.percent.
	XPercent *InputXPercent `json:"xpercent,omitempty"`
}

// InputXPercent is value transformation config for status.percent.
type InputXPercent struct {
	// Expr is a JS expression containing x, where x is the raw value to be transformed
	// (status.percent), for example "x+1". Accepted range: null or [0..100] chars. Both
	// null and "" mean value transformation is disabled.
	Expr *string `json:"expr,omitempty"`

	// Unit of the transformed value (status.xpercent), for example, "m/s".
	// Accepted range: null or [0..20] chars. Both null and "" mean value transformation
	// is disabled.
	Unit *string `json:"unit,omitempty"`
}

type InputSetConfigRequest struct {
	// ID of the input component instance.
	ID int `json:"id"`

	// Configuration that the method takes.
	Config InputConfig `json:"config"`
}

func (r *InputSetConfigRequest) Method() string {
	return "Input.SetConfig"
}

func (r *InputSetConfigRequest) NewResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *InputSetConfigRequest) Do(
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

type InputCheckExpressionRequest struct {
	// Expr is the JS expression to evaluate.
	Expr string `json:"expr,omitempty"`

	// Inputs on which to apply expr. Elements are allowed to be null
	Inputs []*float64 `json:"inputs,omitempty"`
}

func (r *InputCheckExpressionRequest) Method() string {
	return "Input.CheckExpression"
}

func (r *InputCheckExpressionRequest) NewResponse() *InputCheckExpressionResponse {
	return &InputCheckExpressionResponse{}
}

func (r *InputCheckExpressionRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*InputCheckExpressionResponse,
	*frame.Response,
	error,
) {
	resp := r.NewResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

type InputCheckExpressionResponse struct {
	Results []InputCheckExpressionResult
}

type InputCheckExpressionResult struct {
	Input *float64

	Output *float64

	Error *string
}

func (r *InputCheckExpressionResult) UnmarshalJSON(b []byte) error {
	var got []interface{}
	if err := json.Unmarshal(b, &got); err != nil {
		return err
	}
	if len(got) >= 1 {
		if n, ok := got[0].(json.Number); ok {
			f64, err := n.Float64()
			if err != nil {
				return fmt.Errorf("parsing input result: %w", err)
			}
			r.Input = Float64Ptr(f64)
		} else if n, ok := got[0].(float64); ok {
			r.Input = Float64Ptr(n)
		}

	}
	if len(got) >= 2 {
		if n, ok := got[1].(json.Number); ok {
			f64, err := n.Float64()
			if err != nil {
				return fmt.Errorf("parsing output result: %w", err)
			}
			r.Output = Float64Ptr(f64)
		} else if n, ok := got[1].(float64); ok {
			r.Output = Float64Ptr(n)
		}
	}
	if len(got) >= 3 {
		if s, ok := got[2].(string); ok {
			r.Error = StrPtr(s)
		}
	}

	return nil
}
