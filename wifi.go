package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

type WifiStatus struct {
	// StaIP is the IP of the device in the network (null if disconnected).
	StaIP *string `json:"sta_ip,omitempty"`

	// Status of the connection. Range of values: disconnected, connecting, connected, got ip
	Status string `json:"status,omitempty"`

	// SSID of the network (null if disconnected)
	SSID *string `json:"ssid,omitempty"`

	// RSSI is the strength of the signal in dBms.
	RRSI *float64 `json:"rssi,omitempty"`

	// APClientCount is the number of clients connected to the access point. Present only when
	// AP is enabled and range extender functionality is present and enabled.
	APClientCount *int `json:"ap_client_count,omitempty"`
}

type WifiGetStatusRequest struct{}

func (r *WifiGetStatusRequest) Method() string {
	return "Wifi.GetStatus"
}

func (r *WifiGetStatusRequest) NewTypedResponse() *WifiStatus {
	return &WifiStatus{}
}

func (r *WifiGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *WifiGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*WifiStatus,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type WifiConfig struct {
	// AP is information about the access point.
	AP *WifiAPConfig `json:"ap,omitempty"`

	// STA is the wifi station config.
	STA *WifiStationConfig `json:"sta,omitempty"`

	// STA1 will be used as fallback when the device is unable to connect to the sta network.
	STA1 *WifiStationConfig `json:"sta1,omitempty"`

	Roam *WifiRoamConfig `json:"roam,omitempty"`
}

type WifiAPConfig struct {
	// SSID is the readonly SSID of the access point
	SSID *string `json:"ssid,omitempty"`

	// Pass is the password for the ssid, writeonly. Must be provided if you provide ssid.
	Pass *string `json:"pass,omitempty"`

	// IsOpen is true if the access point is open, false otherwise.
	IsOpen *bool `json:"is_open,omitempty"`

	// Enable is true if the access point is enabled, false otherwise
	Enable *bool `json:"enable,omitempty"`

	// RangeExtender configuration object, available only when range extender functionality
	// is present.
	RangeExtender *WifiAPRangeExtenderConfig `json:"range_extender,omitempty"`
}

type WifiAPRangeExtenderConfig struct {
	// Enable is true if range extender functionality is enabled, false otherwise
	Enable *bool `json:"enable,omitempty"`
}

type WifiStationConfig struct {
	// SSID of the network.
	SSID *string `json:"ssid,omitempty"`

	// Pass is the password for the ssid, writeonly. Must be provided if you provide ssid.
	Pass *string `json:"pass,omitempty"`

	// IsOpen is true if the network is open, i.e. no password is set, false otherwise,
	// readonly.
	IsOpen *bool `json:"is_open,omitempty"`

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

type WifiRoamConfig struct {
	// RSSI_Thr is the RSSI threshold - when reached will trigger the access point
	// roaming. Default value: -80
	RSSI_Thr *float64 `json:"rssi_thr,omitempty"`

	// Interval at which to scan for better access points. Enabled if set to positive
	// number, disabled if set to 0. Default value: 60
	Interval *float64 `json:"interval,omitempty"`
}

type WifiGetConfigRequest struct{}

func (r *WifiGetConfigRequest) Method() string {
	return "Wifi.GetConfig"
}

func (r *WifiGetConfigRequest) NewTypedResponse() *WifiConfig {
	return &WifiConfig{}
}

func (r *WifiGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *WifiGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*WifiConfig,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type WifiSetConfigRequest struct {
	Config WifiConfig `json:"config"`
}

func (r *WifiSetConfigRequest) Method() string {
	return "Wifi.SetConfig"
}

func (r *WifiSetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *WifiSetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *WifiSetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*SetConfigResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}
