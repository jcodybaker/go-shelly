package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

type MQTTSetConfigRequest struct {
	Config MQTTConfig `json:"config"`
}

func (r *MQTTSetConfigRequest) Method() string {
	return "MQTT.SetConfig"
}

func (r *MQTTSetConfigRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *MQTTSetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *MQTTSetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*RPCEmptyResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type MQTTGetConfigRequest struct {
	Config MQTTConfig `json:"config"`
}

func (r *MQTTGetConfigRequest) Method() string {
	return "MQTT.GetConfig"
}

func (r *MQTTGetConfigRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *MQTTGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *MQTTGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*RPCEmptyResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

// MQTTConfig configures MQTT for Shelly.
type MQTTConfig struct {
	// Enbable is true if MQTT connection is enabled, false otherwise
	Enable bool `json:"enabled"`
	// Server is the hostname of the MQTT server. Can be followed by port number - host:port
	Server *string `json:"server"`
	// ClientID identifies each MQTT client that connects to an MQTT brokers. Defaults if null to device id.
	ClientID *string `json:"client_id"`
	// User is the username.
	User *string `json:"user"`
	// SSL_CA determines the type of connection to make.
	// If null, no TLS will be used.
	// If `*` TLS connections will be made without server verification.
	// If `user_ca.pem` TLS connection will be verified by the user-provided CA.
	// If `ca.pem` TLS connections will be verified against the default CA list.
	SSL_CA *string `json:"ssl_ca"`
	// TopicPrefix is the prefix of the topics on which device publish/subscribe. Limited to 300
	// characters. Could not start with $ and #, +, %, ? are not allowed.
	TopicPrefix *string `json:"topic_prefix"`
	// RPC_NTF enables RPC notifications (NotifyStatus and NotifyEvent) to be published on
	// <device_id|topic_prefix>/events/rpc (<topic_prefix> when a custom prefix is set, <device_id>
	// otherwise). Default value: true.
	RPC_NTF bool `json:"rpc_ntf"`
	// STATUS_NTF Enables publishing the complete component status on
	// <device_id|topic_prefix>/status/<component>:<id> (<topic_prefix> when a custom prefix is set,
	// <device_id> otherwise). The complete status will be published if a signifficant change
	// occurred. Default value: false
	STATUS_NTF bool `json:"status_ntf"`
	// UseClientCert enables or disables usage of client certifactes to use MQTT with encription,
	// default: false
	UseClientCert bool `json:"use_client_cert"`
	// EnableControl enables the MQTT control feature. Defalut value: true
	EnableControl bool `json:"enable_control"`
}

type MQTTGetStatusRequest struct{}

// Method returns the method name.
func (r *MQTTGetStatusRequest) Method() string {
	return "MQTT.GetStatus"
}

func (r *MQTTGetStatusRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *MQTTGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *MQTTGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*RPCEmptyResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type MQTTStatus struct {
	Connected bool `json:"connected"`
}
