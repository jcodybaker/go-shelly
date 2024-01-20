package shelly

import (
	"context"
	"encoding/json"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

// Schedule describes a series of RPCs to be repeated on a schedule.
type Schedule struct {
	// ID assigned to the job when it is created. This is used in subsequent Update / Delete calls.
	// This should be nil for Schedule.Create, and MUST have a value for Schedule.Update.
	ID *int `json:"id,omitempty"`

	// Enable is true to enable the execution of this job, false otherwise. It is true by default.
	Enable *bool `json:"enable,omitempty"`

	// Timespec as defined by [cron](https://github.com/mongoose-os-libs/cron). Note that leading
	// 0s are not supported (e.g.: for 8 AM you should set 8 instead of 08).
	TimeSpec *string `json:"timespec,omitempty"`

	// Calls is a list of RPC methods and arguments to be invoked when the job gets executed. It
	// must contain at least one valid object. There is a limit of 5 calls per schedule job.
	Calls []ScheduleCall `json:"calls,omitempty"`
}

// ScheduleCall describes a single RPC call to be initiated by the Schedule.
type ScheduleCall struct {
	// Method is the name of the RPC method. Required
	Method string `json:"method"`

	// Params are the parameters used to invoke the RPC call. If the call requires no parameters
	// params may be omitted.
	Params *json.RawMessage `json:"params,omitempty"`
}

// NewScheduleCallRPCRequest creates a ScheduleCall object from a RPCRequestBody.
func NewScheduleCallWithRPCRequest(req RPCRequestBody) (ScheduleCall, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return ScheduleCall{}, err
	}
	asRaw := json.RawMessage(b)
	return ScheduleCall{
		Method: req.Method(),
		Params: &asRaw,
	}, nil
}

// ScheduleCreateRequest adds a new schedule to the shelly device.
type ScheduleCreateRequest Schedule

func (r *ScheduleCreateRequest) Method() string {
	return "Schedule.Create"
}

func (r *ScheduleCreateRequest) NewTypedResponse() *ScheduleCreateResponse {
	return &ScheduleCreateResponse{}
}

func (r *ScheduleCreateRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ScheduleCreateRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*ScheduleCreateResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// ScheduleCreateResponse is the RPC response to the ScheduleCreateRequest.
type ScheduleCreateResponse struct {
	// ID assigned to the scheduled job.
	ID *int `json:"id,omitempty"`

	// Rev is the current revision number of the schedule instances.
	Rev *int `json:"rev,omitempty"`
}

// ScheduleUpdateResponse is the response for Schedule.Update, Schedule.Delete,
// and Schedule.DeleteAll RPC requests.
type ScheduleUpdateResponse struct {
	// Rev is the current revision number of the schedule instances.
	Rev *int `json:"rev,omitempty"`
}

// ScheduleUpdateRequest modifies an existing schedule.
type ScheduleUpdateRequest Schedule

func (r *ScheduleUpdateRequest) Method() string {
	return "Schedule.Update"
}

func (r *ScheduleUpdateRequest) NewTypedResponse() *ScheduleUpdateResponse {
	return &ScheduleUpdateResponse{}
}

func (r *ScheduleUpdateRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ScheduleUpdateRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*ScheduleUpdateResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// ScheduleDeleteRequest deletes an existing schedule.
type ScheduleDeleteRequest struct {
	// ID of the schedule to be deleted. Required.
	ID int `json:"id"`
}

func (r *ScheduleDeleteRequest) Method() string {
	return "Schedule.Delete"
}

func (r *ScheduleDeleteRequest) NewTypedResponse() *ScheduleUpdateResponse {
	return &ScheduleUpdateResponse{}
}

func (r *ScheduleDeleteRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ScheduleDeleteRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*ScheduleUpdateResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

// ScheduleDeleteAllRequest deletes all existing schedules.
type ScheduleDeleteAllRequest struct{}

func (r *ScheduleDeleteAllRequest) Method() string {
	return "Schedule.Delete"
}

func (r *ScheduleDeleteAllRequest) NewTypedResponse() *ScheduleUpdateResponse {
	return &ScheduleUpdateResponse{}
}

func (r *ScheduleDeleteAllRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ScheduleDeleteAllRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*ScheduleUpdateResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}
