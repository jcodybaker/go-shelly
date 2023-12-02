package shelly_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	shelly "github.com/jcodybaker/go-shelly"
	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

func GetCallWithVerify(t *testing.T, req shelly.RPCRequestBody, respBody interface{}) {
	ctx := context.Background()
	c, err := mgrpc.New(ctx, "http://192.168.1.10/rpc", mgrpc.UseHTTPPost())
	require.NoError(t, err)
	defer c.Disconnect(ctx)
	args, err := json.Marshal(req)
	require.NoError(t, err)
	command := &frame.Command{
		Cmd:  req.Method(),
		Args: json.RawMessage(args),
	}
	resp, err := c.Call(ctx, "", command, nil)
	require.NoError(t, err)
	fmt.Println(string(resp.Response))
	require.NoErrorf(
		t,
		json.Unmarshal(resp.Response, &respBody),
		"got resp code: %d (%s) body: %s",
		resp.Status,
		resp.StatusMsg,
		resp.Response,
	)

	// The reencoded JSON *SHOULD* match.
	// NOTE: in practice there seem to be some undocumented fields and inconsistency in what
	// is NULL and what is omited when NULL.
	jsonOut, err := json.Marshal(respBody)
	require.NoError(t, err)
	assert.JSONEq(t, string(resp.Response), string(jsonOut))
}
