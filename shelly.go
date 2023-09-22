package shelly

import (
	"crypto/sha256"
	"encoding/hex"
)

const (
	// DefaultAuthenticationUsername is the only username allowed for auth.
	DefaultAuthenticationUsername = "admin"
)

type ShellyGetStatusRequest struct{}

func (r *ShellyGetStatusRequest) Method() string {
	return "Shelly.GetStatus"
}

type ShellyGetStatusResponse struct {
	MQTT *MQTTStatus `json:"mqtt,omitempty"`

	Cloud *CloudStatus `json:"cloud,omitempty"`
}

type ShellyGetDeviceInfoRequest struct {
	// Ident is a flag specifying if extra identifying information should be displayed.
	Ident bool
}

func (r *ShellyGetDeviceInfoRequest) Method() string {
	return "Shelly.GetDeviceInfo"
}

type ShellyGetDeviceInfoResponse struct {
	// ID of the device.
	ID string `json:"id"`

	// MAC of the device.
	MAC string `json:"mac"`

	// Model of the device
	Model string `json:"model"`

	// Gen is the generation of the device
	Gen string `json:"gen"`

	// FW_ID is the firmware id of the device.
	FW_ID string `json:"fw_id"`

	// Ver is the version of the device firmware.
	Ver string `json:"ver"`

	// App is the application name.
	App string `json:"app"`

	// Profile is the name of the device profile (only applicable for multi-profile devices)
	Profile string `json:"profile"`

	// AuthEn is true if authentication is enabled.
	AuthEn string `json:"auth_en"`

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

type ShellyCheckForUpdateResponse struct {
	// Stable indicates the new stable version of the firmware.
	Stable FirmwareUpdateVersion

	// Beta indicates the new beta version of the firmware.
	Beta FirmwareUpdateVersion
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
	Stage string `json:"stage"`

	// URL address of the update. (Optional). If set Stage must be "".
	URL string `json:"url"`
}

type ShellyFactoryResetRequest struct{}

func (r *ShellyFactoryResetRequest) Method() string {
	return "Shelly.FactoryReset"
}

type ShellyResetWiFiConfigRequest struct{}

func (r *ShellyResetWiFiConfigRequest) Method() string {
	return "Shelly.ResetWiFiConfig"
}

type ShellyRebootRequest struct {
	// DelayMS sets the delay until reboot in milliseconds. Any values are valid but the minimum
	// is capped at 500 ms. Default value: 1000 ms. (Optional)
	DelayMS int `json:"delay_ms,omitempty"`
}

func (r *ShellyRebootRequest) Method() string {
	return "Shelly.Reboot"
}

type ShellySetAuthRequest struct {
	// User must be set to admin. Only one user is supported. (Required)
	User string

	// Realm must be the id of the device. Only one realm is supported. (Required)
	Realm string

	// HA1 "user:realm:password" encoded in SHA256 (null to disable authentication).
	HA1 *string
}

func (r *ShellySetAuthRequest) Method() string {
	return "Shelly.SetAuth"
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
		hex.EncodeToString(
			sha256.New().Sum([]byte(out.User + ":" + out.Realm + ":" + password))))
	return out
}

type ShellyPutUserCARequest struct {
	// Contents of the PEM file (null if you want to delete the existing data). (Required)
	Data *string

	// Append is true if more data will be appended afterwards, default false.
	Append bool
}

func (r *ShellyPutUserCARequest) Method() string {
	return "Shelly.PutUserCA"
}

type ShellyPutTLSClientCertRequest struct {
	// Contents of the PEM file (null if you want to delete the existing data). (Required)
	Data *string

	// Append is true if more data will be appended afterwards, default false.
	Append bool
}

func (r *ShellyPutTLSClientCertRequest) Method() string {
	return "Shelly.PutTLSClientCert"
}

type ShellyPutTLSClientKeyRequest struct {
	// Contents of the PEM file (null if you want to delete the existing data). (Required)
	Data *string

	// Append is true if more data will be appended afterwards, default false.
	Append bool
}

func (r *ShellyPutTLSClientKeyRequest) Method() string {
	return "Shelly.PutTLSClientKey"
}
