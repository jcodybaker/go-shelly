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
	c, err := mgrpc.New(ctx, "http://192.168.1.62/rpc", mgrpc.UseHTTPPost())
	require.NoError(t, err)
	defer c.Disconnect(ctx)
	command := &frame.Command{
		Cmd: req.Method(),
	}
	resp, err := c.Call(ctx, "", command, nil)
	require.NoError(t, err)
	fmt.Println(string(resp.Response))
	require.NoError(t, json.Unmarshal(resp.Response, &respBody))

	// The reencoded JSON *SHOULD* match.
	// NOTE: in practice there seem to be some undocumented fields and inconsistency in what
	// is NULL and what is omited when NULL.
	jsonOut, err := json.Marshal(respBody)
	require.NoError(t, err)
	assert.JSONEq(t, string(resp.Response), string(jsonOut))
}
