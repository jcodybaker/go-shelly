package shelly

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

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
		return rawResp, &BadStatusWithMessageError{Status: ShellyErrorCode(rawResp.Status), Msg: rawResp.StatusMsg}
	}
	if err := json.Unmarshal(rawResp.Response, resp); err != nil {
		return rawResp, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return rawResp, err
}
