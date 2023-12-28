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
)

func GetCallWithVerify(t *testing.T, req shelly.RPCRequestBody, respBody interface{}) {
	ctx := context.Background()
	c, err := mgrpc.New(ctx, "http://192.168.1.23/rpc", mgrpc.UseHTTPPost())
	require.NoError(t, err)
	defer c.Disconnect(ctx)

	respFrame, err := shelly.Do(ctx, c, req, respBody)
	require.NoError(t, err)
	fmt.Println(string(respFrame.Response))

	// The reencoded JSON *SHOULD* match.
	// NOTE: in practice there seem to be some undocumented fields and inconsistency in what
	// is NULL and what is omited when NULL.
	jsonOut, err := json.Marshal(respBody)
	require.NoError(t, err)
	assert.JSONEq(t, string(respFrame.Response), string(jsonOut))
}
