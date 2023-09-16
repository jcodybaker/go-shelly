package shelly

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

const (
	JSON_RPC_VERSION = "2.0"
)

// RPCRequest is the outermost framing for calls to the shelly device.
type RPCRequest struct {
	// ID sets a unique identifier for matching responses with requests. (Required)
	ID string

	// Source sets the name of the source of the request. This may be an arbitary string. (Required)
	Source string

	// Body (Required)
	Body RPCRequestBody
}

// Raw converts the request to a RawRPCRequest ready for the wire.
func (r *RPCRequest) Raw() (*RawRPCRequest, error) {
	if r.Body == nil {
		return nil, errors.New("RPCRequest body is required")
	}
	params := new(bytes.Buffer)
	if err := json.NewEncoder(params).Encode(r.Body); err != nil {
		return nil, fmt.Errorf("encoding request body: %w", err)
	}
	return &RawRPCRequest{
		ID:         r.ID,
		Source:     r.Source,
		JSON_RPC:   JSON_RPC_VERSION,
		Method:     r.Body.Method(),
		Parameters: params.Bytes(),
	}, nil
}

// InitID sets an randomoized ID on the RPCRequest.
func (r *RPCRequest) InitID() error {
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("building rpc request id: %w", err)
	}
	r.ID = id.String()
	return nil
}

// RawRPCRequest is the outermost framing for calls to the shelly device.
type RawRPCRequest struct {
	// JSON_RPC describes the JSON RPC version used. This library implements 2.0.
	JSON_RPC string `json:"jsonrpc,omitempty"`

	// ID sets a unique identifier for matching responses with requests. (Required)
	ID string `json:"id,omitempty"`

	// Source sets the name of the source of the request. This may be an arbitary string. (Required)
	Source string `json:"source,omitempty"`

	// Method is the name of the proceedure to be called (Required).
	Method string `json:"method,omitempty"`

	// Parameters is the key-value dictionary of parameters for the Method, if any.
	Parameters json.RawMessage `json:"params,omitempty"`
}

// RPCRequestBody describes objects implementing an RPC request body.
type RPCRequestBody interface {
	Method() string
}

type RawRPCResponse struct {
	// ID sets a unique identifier for matching responses with requests. (Required)
	ID string `json:"id,omitempty"`

	//Name of the source of the response (the Shelly device). (Required)
	Source string `json:"source,omitempty"`

	// Destination describes the original source of the request. (Required)
	Destination string `json:"dst,omitempty"`

	Result json.RawMessage

	Error *RPCError
}

type RPCResponse struct {
}

// RawRPCNotification is the outermost framing for notifications
type RawRPCNotification struct {
	// JSON_RPC describes the JSON RPC version used. This library implements 2.0.
	JSON_RPC string `json:"jsonrpc,omitempty"`

	// Source sets the name of the source of the request. This may be an arbitary string. (Required)
	Source string `json:"source,omitempty"`

	// Method is the name of the proceedure to be called (Required).
	Method string `json:"method,omitempty"`

	// Parameters is the key-value dictionary of parameters for the Method, if any.
	Parameters json.RawMessage `json:"params,omitempty"`
}
