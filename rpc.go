package shelly

// RPCRequestBody describes objects implementing an RPC request body.
type RPCRequestBody interface {
	Method() string
}
