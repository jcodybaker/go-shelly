package shelly

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

const (
	// DefaultAuthenticationUsername is the only username allowed for auth.
	DefaultAuthenticationUsername = "admin"
)

type ShellyGetStatusRequest struct{}

func (r *ShellyGetStatusRequest) Method() string {
	return "Shelly.GetStatus"
}

func (r *ShellyGetStatusRequest) NewTypedResponse() *ShellyGetStatusResponse {
	return &ShellyGetStatusResponse{}
}

func (r *ShellyGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*ShellyGetStatusResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type ShellyGetStatusResponse struct {
	System *SysStatus `json:"sys,omitempty"`

	Wifi *WifiStatus `json:"wifi,omitempty"`

	Ethernet *EthStatus `json:"eth,omitempty"`

	BLE *BLEStatus `json:"ble,omitempty"`

	Cloud *CloudStatus `json:"cloud,omitempty"`

	MQTT *MQTTStatus `json:"mqtt,omitempty"`

	// WebSocket *WsStatus `json:"ws,omitempty"`

	// Scripts []*ScriptStatus

	Inputs []*InputStatus `json:"inputs,omitempty"`

	// ModBus *ModBusStatus

	// Voltmeters []*VoltmeterStatus

	Covers []*CoverStatus `json:"covers,omitempty"`

	Switches []*SwitchStatus `json:"switches,omitempty"`

	Lights []*LightStatus `json:"lights,omitempty"`

	DevicePowers []*DevicePowerStatus `json:"device_powers,omitempty"`

	Humidities []*HumidityStatus `json:"humidities,omitempty"`

	Temperatures []*TemperatureStatus `json:"temperatures,omitempty"`

	// EMs []*EMStatus

	// EM1s []*EM1Status

	// PM1s []*PM1Status

	// EMDatas []*EMDataStatus

	// EM1Datas []EM1DataStatus

	// Smokes []*SmokeStatus
}

func (r *ShellyGetStatusResponse) UnmarshalJSON(b []byte) error {
	theRest := make(map[string]json.RawMessage)
	if err := json.Unmarshal(b, &theRest); err != nil {
		return err
	}
	if v, ok := theRest["ble"]; ok {
		var s BLEStatus
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.BLE = &s
	}
	if v, ok := theRest["sys"]; ok {
		var s SysStatus
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.System = &s
	}
	if v, ok := theRest["cloud"]; ok {
		var s CloudStatus
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Cloud = &s
	}
	if v, ok := theRest["mqtt"]; ok {
		var s MQTTStatus
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.MQTT = &s
	}
	if v, ok := theRest["wifi"]; ok {
		var s WifiStatus
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Wifi = &s
	}
	if v, ok := theRest["eth"]; ok {
		var s EthStatus
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Ethernet = &s
	}

	for i := 0; ; i++ {
		v, ok := theRest[fmt.Sprintf("switch:%d", i)]
		if !ok {
			break
		}
		var s SwitchStatus
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Switches = append(r.Switches, &s)
	}
	for i := 0; ; i++ {
		v, ok := theRest[fmt.Sprintf("cover:%d", i)]
		if !ok {
			break
		}
		var s CoverStatus
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Covers = append(r.Covers, &s)
	}
	for i := 0; ; i++ {
		v, ok := theRest[fmt.Sprintf("input:%d", i)]
		if !ok {
			break
		}
		var s InputStatus
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Inputs = append(r.Inputs, &s)
	}
	for i := 0; ; i++ {
		v, ok := theRest[fmt.Sprintf("light:%d", i)]
		if !ok {
			break
		}
		var s LightStatus
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Lights = append(r.Lights, &s)
	}
	for i := 0; ; i++ {
		v, ok := theRest[fmt.Sprintf("devicepower:%d", i)]
		if !ok {
			break
		}
		var s DevicePowerStatus
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.DevicePowers = append(r.DevicePowers, &s)
	}
	for i := 0; ; i++ {
		v, ok := theRest[fmt.Sprintf("humidity:%d", i)]
		if !ok {
			break
		}
		var s HumidityStatus
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Humidities = append(r.Humidities, &s)
	}
	for i := 0; ; i++ {
		v, ok := theRest[fmt.Sprintf("temperature:%d", i)]
		if !ok {
			break
		}
		var s TemperatureStatus
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Temperatures = append(r.Temperatures, &s)
	}
	return nil
}

type ShellyGetDeviceInfoRequest struct {
	// Ident is a flag specifying if extra identifying information should be displayed.
	Ident bool
}

func (r *ShellyGetDeviceInfoRequest) Method() string {
	return "Shelly.GetDeviceInfo"
}

func (r *ShellyGetDeviceInfoRequest) NewTypedResponse() *ShellyGetDeviceInfoResponse {
	return &ShellyGetDeviceInfoResponse{}
}

func (r *ShellyGetDeviceInfoRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyGetDeviceInfoRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*ShellyGetDeviceInfoResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type ShellyGetDeviceInfoResponse struct {
	// ID of the device.
	ID string `json:"id"`

	// MAC of the device.
	MAC string `json:"mac"`

	// Model of the device
	Model string `json:"model"`

	// Gen is the generation of the device
	Gen json.Number `json:"gen"`

	// FW_ID is the firmware id of the device.
	FW_ID string `json:"fw_id"`

	// Ver is the version of the device firmware.
	Ver string `json:"ver"`

	// App is the application name.
	App string `json:"app"`

	// Profile is the name of the device profile (only applicable for multi-profile devices)
	Profile string `json:"profile"`

	// AuthEn is true if authentication is enabled.
	AuthEn bool `json:"auth_en"`

	// Name of the domain (null if authentication is not enabled)
	AuthDomain *string `json:"auth_domain"`

	// Present only when false. If true, device is shown in 'Discovered devices'. If false, the device is hidden.
	Discoverable *bool `json:"discoverable"`

	// Key is cloud key of the device (see note below), present only when the ident parameter is set to true.
	Key string `json:"key"`

	// Batch used to provision the device, present only when the ident parameter is set to true.
	Batch string `json:"batch"`

	// FW_SBits are shelly internal flags, present only when the ident parameter is set to true.
	FW_SBits string
}

// ShellyCheckForUpdateRequest
type ShellyCheckForUpdateRequest struct{}

func (r *ShellyCheckForUpdateRequest) Method() string {
	return "Shelly.CheckForUpdate"
}

func (r *ShellyCheckForUpdateRequest) NewTypedResponse() *ShellyCheckForUpdateResponse {
	return &ShellyCheckForUpdateResponse{}
}

func (r *ShellyCheckForUpdateRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyCheckForUpdateRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*ShellyCheckForUpdateResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type ShellyCheckForUpdateResponse struct {
	// Stable indicates the new stable version of the firmware.
	Stable *FirmwareUpdateVersion `json:"stable,omitempty"`

	// Beta indicates the new beta version of the firmware.
	Beta *FirmwareUpdateVersion `json:"beta,omitempty"`
}

type FirmwareUpdateVersion struct {
	// Version is the new version available.
	Version string `json:"version,omitempty"`

	// BuildID is the build ID of the update.
	BuildID string `json:"build_id,omitempty"`
}

type ShellyUpdateRequest struct {
	// Stage is the the type of the new version - either stable or beta. By default updates to
	// stable version. (Optional)
	Stage string `json:"stage,omitempty"`

	// URL address of the update. (Optional). If set Stage must be "".
	URL string `json:"url,omitempty"`
}

func (r *ShellyUpdateRequest) Method() string {
	return "Shelly.Update"
}

func (r *ShellyUpdateRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *ShellyUpdateRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type ShellyFactoryResetRequest struct{}

func (r *ShellyFactoryResetRequest) Method() string {
	return "Shelly.FactoryReset"
}

func (r *ShellyFactoryResetRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *ShellyFactoryResetRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyFactoryResetRequest) Do(
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

type ShellyResetWiFiConfigRequest struct{}

func (r *ShellyResetWiFiConfigRequest) Method() string {
	return "Shelly.ResetWiFiConfig"
}

func (r *ShellyResetWiFiConfigRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *ShellyResetWiFiConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyResetWiFiConfigRequest) Do(
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

type ShellyRebootRequest struct {
	// DelayMS sets the delay until reboot in milliseconds. Any values are valid but the minimum
	// is capped at 500 ms. Default value: 1000 ms. (Optional)
	DelayMS int `json:"delay_ms,omitempty"`
}

func (r *ShellyRebootRequest) Method() string {
	return "Shelly.Reboot"
}

func (r *ShellyRebootRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *ShellyRebootRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyRebootRequest) Do(
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

type ShellySetAuthRequest struct {
	// User must be set to admin. Only one user is supported. (Required)
	User string `json:"user,omitempty"`

	// Realm must be the id of the device. Only one realm is supported. (Required)
	Realm string `json:"realm,omitempty"`

	// HA1 "user:realm:password" encoded in SHA256 (null to disable authentication).
	HA1 *string `json:"ha1"`
}

func (r *ShellySetAuthRequest) Method() string {
	return "Shelly.SetAuth"
}

func (r *ShellySetAuthRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *ShellySetAuthRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellySetAuthRequest) Do(
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

func NewShellySetAuthRequest(deviceID, password string) *ShellySetAuthRequest {
	out := &ShellySetAuthRequest{
		User:  DefaultAuthenticationUsername,
		Realm: deviceID,
	}
	if password == "" {
		return out
	}
	out.HA1 = StrPtr(
		fmt.Sprintf("%x",
			sha256.Sum256([]byte(out.User+":"+out.Realm+":"+password))))
	return out
}

// BuildShellyAuthRequest builds the request, fetching the deviceID for realm.
func BuildShellyAuthRequest(
	ctx context.Context,
	c mgrpc.MgRPC,
	password string,
) (*ShellySetAuthRequest, error) {
	resp, _, err := (&ShellyGetDeviceInfoRequest{}).Do(ctx, c, nil)
	if err != nil {
		return nil, err
	}
	deviceID := resp.ID
	return NewShellySetAuthRequest(deviceID, password), nil
}

type ShellyPutUserCARequest struct {
	// Contents of the PEM file (null if you want to delete the existing data). (Required)
	Data *string `json:"data"`

	// Append is true if more data will be appended afterwards, default false.
	Append bool `json:"append,omitempty"`
}

func (r *ShellyPutUserCARequest) Method() string {
	return "Shelly.PutUserCA"
}

func (r *ShellyPutUserCARequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *ShellyPutUserCARequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyPutUserCARequest) Do(
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

// ShellyPutUserCA is a helper method which uploads the provided data to the Shelly.PutUserCA method,
// line-by-line to accomodate limits on payload size.
func ShellyPutUserCA(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
	data io.Reader,
) error {
	s := bufio.NewScanner(data)
	req := &ShellyPutUserCARequest{}
	for s.Scan() {
		req.Data = StrPtr(s.Text() + "\n")
		if _, _, err := req.Do(ctx, c, credsCallback); err != nil {
			return err
		}
		req.Append = true
	}
	if err := s.Err(); err != nil {
		return fmt.Errorf("reading input data for Shelly.PutUserCA: %w", err)
	}
	return nil
}

type ShellyPutTLSClientCertRequest struct {
	// Data contents of the PEM file (null if you want to delete the existing data). (Required)
	Data *string `json:"data"`

	// Append is true if more data will be appended afterwards, default false.
	Append bool `json:"append,omitempty"`
}

func (r *ShellyPutTLSClientCertRequest) Method() string {
	return "Shelly.PutTLSClientCert"
}

func (r *ShellyPutTLSClientCertRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *ShellyPutTLSClientCertRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyPutTLSClientCertRequest) Do(
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

// ShellyPutTLSClientCert is a helper method which uploads the provided data to the
// Shelly.PutTLSClientCert method, line-by-line to accomodate limits on payload size.
func ShellyPutTLSClientCert(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
	data io.Reader,
) error {
	s := bufio.NewScanner(data)
	req := &ShellyPutTLSClientCertRequest{}
	for s.Scan() {
		req.Data = StrPtr(s.Text() + "\n")
		if _, _, err := req.Do(ctx, c, credsCallback); err != nil {
			return err
		}
		req.Append = true
	}
	if err := s.Err(); err != nil {
		return fmt.Errorf("reading input data for Shelly.TLSClientCert: %w", err)
	}
	return nil
}

type ShellyPutTLSClientKeyRequest struct {
	// Contents of the PEM file (null if you want to delete the existing data). (Required)
	Data *string `json:"data"`

	// Append is true if more data will be appended afterwards, default false.
	Append bool `json:"append,omitempty"`
}

func (r *ShellyPutTLSClientKeyRequest) Method() string {
	return "Shelly.PutTLSClientKey"
}

func (r *ShellyPutTLSClientKeyRequest) NewTypedResponse() *RPCEmptyResponse {
	return &RPCEmptyResponse{}
}

func (r *ShellyPutTLSClientKeyRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyPutTLSClientKeyRequest) Do(
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

// ShellyPutTLSClientKey is a helper method which uploads the provided data to the
// Shelly.PutTLSClientKey method, line-by-line to accomodate limits on payload size.
func ShellyPutTLSClientKey(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
	data io.Reader,
) error {
	s := bufio.NewScanner(data)
	req := &ShellyPutTLSClientKeyRequest{}
	for s.Scan() {
		req.Data = StrPtr(s.Text() + "\n")
		if _, _, err := req.Do(ctx, c, credsCallback); err != nil {
			return err
		}
		req.Append = true
	}
	if err := s.Err(); err != nil {
		return fmt.Errorf("reading input data for Shelly.PutTLSClientKey: %w", err)
	}
	return nil
}

type ShellyGetConfigResponse struct {
	System *SysConfig `json:"sys,omitempty"`

	Wifi *WifiConfig `json:"wifi,omitempty"`

	Ethernet *EthConfig `json:"eth,omitempty"`

	BLE *BLEConfig `json:"ble,omitempty"`

	Cloud *CloudConfig `json:"cloud,omitempty"`

	MQTT *MQTTConfig `json:"mqtt,omitempty"`

	// WebSocket *WsConfig `json:"ws,omitempty"`

	// Scripts []*ScriptConfig

	Inputs []*InputConfig `json:"inputs,omitempty"`

	// ModBus *ModBusConfig

	// Voltmeters []*VoltmeterConfig

	Covers []*CoverConfig `json:"covers,omitempty"`

	Switches []*SwitchConfig `json:"switches,omitempty"`

	Lights []*LightConfig `json:"lights,omitempty"`

	// DevicePowers []*DevicePowerConfig

	Humidities []*HumidityConfig `json:"humidities,omitempty"`

	Temperatures []*TemperatureConfig `json:"temperatures,omitempty"`

	// EMs []*EMConfig

	// EM1s []*EM1Config

	// PM1s []*PM1Config

	// EMDatas []*EMDataConfig

	// EM1Datas []EM1DataConfig

	// Smokes []*SmokeConfig
}

func (r *ShellyGetConfigResponse) UnmarshalJSON(b []byte) error {
	theRest := make(map[string]json.RawMessage)
	if err := json.Unmarshal(b, &theRest); err != nil {
		return err
	}
	if v, ok := theRest["ble"]; ok {
		var s BLEConfig
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.BLE = &s
	}
	if v, ok := theRest["sys"]; ok {
		var s SysConfig
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.System = &s
	}
	if v, ok := theRest["cloud"]; ok {
		var s CloudConfig
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Cloud = &s
	}
	if v, ok := theRest["mqtt"]; ok {
		var s MQTTConfig
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.MQTT = &s
	}
	if v, ok := theRest["wifi"]; ok {
		var s WifiConfig
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Wifi = &s
	}
	if v, ok := theRest["eth"]; ok {
		var s EthConfig
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Ethernet = &s
	}

	for i := 0; ; i++ {
		v, ok := theRest[fmt.Sprintf("switch:%d", i)]
		if !ok {
			break
		}
		var s SwitchConfig
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Switches = append(r.Switches, &s)
	}
	for i := 0; ; i++ {
		v, ok := theRest[fmt.Sprintf("cover:%d", i)]
		if !ok {
			break
		}
		var s CoverConfig
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Covers = append(r.Covers, &s)
	}
	for i := 0; ; i++ {
		v, ok := theRest[fmt.Sprintf("input:%d", i)]
		if !ok {
			break
		}
		var s InputConfig
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Inputs = append(r.Inputs, &s)
	}
	for i := 0; ; i++ {
		v, ok := theRest[fmt.Sprintf("light:%d", i)]
		if !ok {
			break
		}
		var s LightConfig
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Lights = append(r.Lights, &s)
	}
	for i := 0; ; i++ {
		v, ok := theRest[fmt.Sprintf("humidity:%d", i)]
		if !ok {
			break
		}
		var s HumidityConfig
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Humidities = append(r.Humidities, &s)
	}
	for i := 0; ; i++ {
		v, ok := theRest[fmt.Sprintf("temperature:%d", i)]
		if !ok {
			break
		}
		var s TemperatureConfig
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		r.Temperatures = append(r.Temperatures, &s)
	}
	return nil
}

type ShellyGetConfigRequest struct{}

func (r *ShellyGetConfigRequest) Method() string {
	return "Shelly.GetConfig"
}

func (r *ShellyGetConfigRequest) NewTypedResponse() *ShellyGetConfigResponse {
	return &ShellyGetConfigResponse{}
}

func (r *ShellyGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*ShellyGetConfigResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type ShellyListMethodsResponse struct {
	// Methods is a list of allowed methods.
	Methods []string `json:"methods,omitempty"`
}

type ShellyListMethodsRequest struct{}

func (r *ShellyListMethodsRequest) Method() string {
	return "Shelly.ListMethods"
}

func (r *ShellyListMethodsRequest) NewTypedResponse() *ShellyListMethodsResponse {
	return &ShellyListMethodsResponse{}
}

func (r *ShellyListMethodsRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyListMethodsRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*ShellyListMethodsResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type ShellyListProfilesComponent struct {
	// Type of component.
	Type string `json:"type"`
	// Count of component instances.
	Count int `json:"count"`
}

type ShellyListProfilesResponse struct {
	// Profiles is a KV mapping of available profiles to an enumeration of their components.
	Profiles map[string][]ShellyListProfilesComponent `json:"profiles,omitempty"`
}

type ShellyListProfilesRequest struct{}

func (r *ShellyListProfilesRequest) Method() string {
	return "Shelly.ListProfiles"
}

func (r *ShellyListProfilesRequest) NewTypedResponse() *ShellyListProfilesResponse {
	return &ShellyListProfilesResponse{}
}

func (r *ShellyListProfilesRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyListProfilesRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*ShellyListProfilesResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type ShellySetProfileResponse struct {
	// ProfileWas will be set to the profile before updating. If the new value is different
	// than the profile_was value, the device will immediately reboot.
	ProfileWas string `json:"profile_was,omitempty"`
}

type ShellySetProfileRequest struct {
	// Profile is the name of the profile to set.
	Profile string `json:"profile"`
}

func (r *ShellySetProfileRequest) Method() string {
	return "Shelly.SetProfile"
}

func (r *ShellySetProfileRequest) NewTypedResponse() *ShellySetProfileResponse {
	return &ShellySetProfileResponse{}
}

func (r *ShellySetProfileRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellySetProfileRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*ShellySetProfileResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type ShellyListTimezonesResponse struct {
	// Timezones is a list of available timezones.
	Timezones []string `json:"timezones"`
}

type ShellyListTimezonesRequest struct{}

func (r *ShellyListTimezonesRequest) Method() string {
	return "Shelly.ListTimezones"
}

func (r *ShellyListTimezonesRequest) NewTypedResponse() *ShellyListTimezonesResponse {
	return &ShellyListTimezonesResponse{}
}

func (r *ShellyListTimezonesRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyListTimezonesRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*ShellyListTimezonesResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type ShellyDetectLocationResponse struct {
	// TZ is the timezone of the detected location (null if not available)
	TZ *string `json:"tz,omitempty"`

	// Lat is the latitude of the detected location in degrees (null if not available)
	Lat *float64 `json:"lat,omitempty"`

	// Lon is the longitude of the detected location in degrees (null if not available)
	Lon *float64 `json:"lon,omitempty"`
}

type ShellyDetectLocationRequest struct{}

func (r *ShellyDetectLocationRequest) Method() string {
	return "Shelly.DetectLocation"
}

func (r *ShellyDetectLocationRequest) NewTypedResponse() *ShellyDetectLocationResponse {
	return &ShellyDetectLocationResponse{}
}

func (r *ShellyDetectLocationRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyDetectLocationRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*ShellyDetectLocationResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

type ShellyComponent struct {
	// Key (in format <type>:<cid>, for example boolean:200)
	Key string `json:"key"`

	// Status, will be omitted if "status" is not specified in the include property.
	Status map[string]interface{}

	// Config, will be omitted if "config" is not specified in the include property.
	Config map[string]interface{}
}

type ShellyGetComponentsResponse struct {
	// Components is a list of ShellyComponent objects.
	Components []*ShellyComponent `json:"components,omitempty"`

	// CfgRev is the configuration revision number from Sys.
	CfgRev int `json:"cfg_rev,omitempty"`

	// Offset is the index of the first component in the result.
	Offset int `json:"offset"`

	// Total number of components with all filters applied.
	Total int `json:"total"`
}

type ShellyGetComponentsRequest struct {
	// Offset is the dndex of the component from which to start generating the result (Optional).
	Offset *int `json:"offset,omitempty"`

	// Include which properties of the component:
	// - "status" will include the component's status
	// - "config" - the config.
	// The keys are always included. Combination of both (["config", "status"]) to get the full
	// config and status of each component. (Optional).
	Include []string `json:"include,omitempty"`

	// DynamicOnly can be set true to include only dynamic components, default false. (Optional)
	DynamicOnly *bool `json:"dynamic_only,omitempty"`
}

func (r *ShellyGetComponentsRequest) Method() string {
	return "Shelly.GetComponents"
}

func (r *ShellyGetComponentsRequest) NewTypedResponse() *ShellyGetComponentsResponse {
	return &ShellyGetComponentsResponse{}
}

func (r *ShellyGetComponentsRequest) NewResponse() any {
	return r.NewTypedResponse()
}

func (r *ShellyGetComponentsRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*ShellyGetComponentsResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, credsCallback, r, resp)
	return resp, raw, err
}

func (r *ShellyGetComponentsRequest) DoAll(
	ctx context.Context,
	c mgrpc.MgRPC,
	credsCallback mgrpc.GetCredsCallback,
) (
	*ShellyGetComponentsResponse,
	error,
) {
	total := 1
	composed := r.NewTypedResponse()
	for have := 0; have < total; {
		resp := r.NewTypedResponse()
		_, err := Do(ctx, c, credsCallback, r, resp)
		if err != nil {
			return nil, err
		}
		composed.Components = append(composed.Components, resp.Components...)
		total = resp.Total
		have += len(resp.Components)
		composed.CfgRev = resp.CfgRev
	}
	return composed, nil
}
