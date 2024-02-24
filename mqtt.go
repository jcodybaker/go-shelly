package shelly

import (
	"context"
	"errors"
	"strings"

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

func (r *MQTTGetConfigRequest) NewTypedResponse() *MQTTConfig {
	return &MQTTConfig{}
}

func (r *MQTTGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *MQTTGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*MQTTConfig,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

// MQTT_SSL_CA is a type to differentiate between not-set (empty string), null (no TLS), and string
// values.
type MQTT_SSL_CA string

const (
	// MQTT_SSL_CA_NULL will disable the TLS on the MQTT connection
	MQTT_SSL_CA_NULL MQTT_SSL_CA = "null"

	// MQTT_SSL_CA_NOT_SET will not send a value for the `ssl_ca` property, leaving it unchanged.
	MQTT_SSL_CA_NOT_SET MQTT_SSL_CA = ""

	// MQTT_SSL_CA_NO_VERIFY will enable TLS but CA skip verification of the server certificate.
	MQTT_SSL_CA_NO_VERIFY MQTT_SSL_CA = "*"

	// MQTT_SSL_CA_DEFAULT_CA will enable TLS with server verification against the default CA bundle.
	MQTT_SSL_CA_DEFAULT_CA MQTT_SSL_CA = "ca.pem"

	// MQTT_SSL_CA_USER_CA will enable TLS with server verification against the user-provided CA.
	// See `Shelly.PutUserCA`.
	MQTT_SSL_CA_USER_CA MQTT_SSL_CA = "user_ca.pem"
)

func (ca *MQTT_SSL_CA) UnmarshalJSON(b []byte) error {
	// NOTE if the balue is unset, this UnmarshallJSON method will not be called which is why
	// MQTT_SSL_CA_NOT_SET is absent.
	s := strings.TrimSpace(string(b))
	switch s {
	case string(MQTT_SSL_CA_NULL):
		*ca = MQTT_SSL_CA_NULL
		return nil
	case `"` + string(MQTT_SSL_CA_NO_VERIFY) + `"`:
		*ca = MQTT_SSL_CA_NO_VERIFY
		return nil
	case `"` + string(MQTT_SSL_CA_DEFAULT_CA) + `"`:
		*ca = MQTT_SSL_CA_DEFAULT_CA
		return nil
	case `"` + string(MQTT_SSL_CA_USER_CA) + `"`:
		*ca = MQTT_SSL_CA_USER_CA
		return nil
	default:
		return errors.New("unknown value for MQTTConfig.SSL_CA")
	}
}

func (ca *MQTT_SSL_CA) MarshalJSON() ([]byte, error) {
	if ca == nil || *ca == MQTT_SSL_CA_NULL {
		return []byte("null"), nil
	}
	return []byte(`"` + *ca + `"`), nil
}

// MQTTConfig configures MQTT for Shelly.
type MQTTConfig struct {
	// Enbable is true if MQTT connection is enabled, false otherwise
	Enable *bool `json:"enabled,omitempty"`
	// Server is the hostname of the MQTT server. Can be followed by port number - host:port
	Server *NullString `json:"server,omitempty"`
	// ClientID identifies each MQTT client that connects to an MQTT brokers. Defaults if null to device id.
	ClientID *NullString `json:"client_id,omitempty"`
	// User is the username.
	User *string `json:"user,omitempty"`
	// Pass is the password.
	Pass *NullString `json:"pass,omitempty"`
	// SSL_CA determines the type of connection to make.
	// If null, no TLS will be used.
	// If `*` TLS connections will be made without server verification.
	// If `user_ca.pem` TLS connection will be verified by the user-provided CA.
	// If `ca.pem` TLS connections will be verified against the default CA list.
	SSL_CA MQTT_SSL_CA `json:"ssl_ca,omitempty"`
	// TopicPrefix is the prefix of the topics on which device publish/subscribe. Limited to 300
	// characters. Could not start with $ and #, +, %, ? are not allowed.
	TopicPrefix *NullString `json:"topic_prefix,omitempty"`
	// RPC_NTF enables RPC notifications (NotifyStatus and NotifyEvent) to be published on
	// <device_id|topic_prefix>/events/rpc (<topic_prefix> when a custom prefix is set, <device_id>
	// otherwise). Default value: true.
	RPC_NTF *bool `json:"rpc_ntf,omitempty"`
	// Status_NTF Enables publishing the complete component status on
	// <device_id|topic_prefix>/status/<component>:<id> (<topic_prefix> when a custom prefix is set,
	// <device_id> otherwise). The complete status will be published if a signifficant change
	// occurred. Default value: false
	Status_NTF *bool `json:"status_ntf,omitempty"`
	// UseClientCert enables or disables usage of client certifactes to use MQTT with encription,
	// default: false
	UseClientCert *bool `json:"use_client_cert,omitempty"`
	// EnableControl enables the MQTT control feature. Defalut value: true
	EnableControl *bool `json:"enable_control,omitempty"`
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
