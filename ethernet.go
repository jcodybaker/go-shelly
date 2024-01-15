package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

type EthStatus struct {
	// IP of the device in the network.
	IP *string `json:"ip"`
}

type EthGetStatusRequest struct{}

func (r *EthGetStatusRequest) Method() string {
	return "Eth.GetStatus"
}

func (r *EthGetStatusRequest) NewResponse() *EthStatus {
	return &EthStatus{}
}

func (r *EthGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*EthStatus,
	*frame.Response,
	error,
) {
	resp := r.NewResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

type EthConfig struct {
	// Enable is true if the configuration is enabled, false otherwise.
	Enable *bool `json:"enable,omitempty"`

	// IPv4Mode Range of values: dhcp, static
	IPv4Mode *string `json:"ipv4mode,omitempty"`

	// IP to use when ipv4mode is static.
	IP *string `json:"ip,omitempty"`

	// Netmask to use when ipv4mode is static
	Netmask *string `json:"netmask,omitempty"`

	// GW is the gateway to use when ipv4mode is static
	GW *string `json:"gw,omitempty"`

	// Nameserver to use when ipv4mode is static
	Nameserver *string `json:"nameserver,omitempty"`
}

type EthGetConfigRequest struct{}

func (r *EthGetConfigRequest) Method() string {
	return "Eth.GetConfig"
}

func (r *EthGetConfigRequest) NewResponse() *EthConfig {
	return &EthConfig{}
}

func (r *EthGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*EthConfig,
	*frame.Response,
	error,
) {
	resp := r.NewResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

type EthSetConfigRequest struct {
	Config EthConfig `json:"config"`
}

func (r *EthSetConfigRequest) Method() string {
	return "Eth.SetConfig"
}

func (r *EthSetConfigRequest) NewResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *EthSetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*SetConfigResponse,
	*frame.Response,
	error,
) {
	resp := r.NewResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}
