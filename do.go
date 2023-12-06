package shelly

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

type BadStatusError struct {
	Status int
	Msg    string
}

func (err *BadStatusError) Error() string {
	return fmt.Sprintf("RPC Bad Status %d: %s", err.Status, err.Msg)
}

func Do[I RPCRequestBody, O any](
	ctx context.Context,
	c mgrpc.MgRPC,
	req I,
	resp O,
) (*frame.Response, error) {
	args, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshalling shelly rpc request: %w", err)
	}
	command := &frame.Command{
		Cmd:  req.Method(),
		Args: json.RawMessage(args),
	}
	rawResp, err := c.Call(ctx, "", command, nil)
	if err != nil {
		return rawResp, fmt.Errorf("making shelly rpc request: %w", err)
	}
	if rawResp.Status != 0 {
		return rawResp, &BadStatusError{Status: rawResp.Status, Msg: rawResp.StatusMsg}
	}
	if err := json.Unmarshal(rawResp.Response, resp); err != nil {
		return rawResp, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return rawResp, err
}
