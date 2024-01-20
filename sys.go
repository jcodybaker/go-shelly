package shelly

import (
	"context"
	"encoding/json"

	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

type SysGetConfigRequest struct{}

func (r *SysGetConfigRequest) Method() string {
	return "Sys.GetConfig"
}

func (r *SysGetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*SysConfig,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

func (r *SysGetConfigRequest) NewTypedResponse() *SysConfig {
	return &SysConfig{}
}

func (r *SysGetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type SysSetConfigRequest struct {
	// Config sets the updated system config.
	Config SysConfig `json:"config"`
}

func (r *SysSetConfigRequest) Method() string {
	return "Sys.SetConfig"
}

func (r *SysSetConfigRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*RPCEmptyResponse,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

func (r *SysSetConfigRequest) NewTypedResponse() *SetConfigResponse {
	return &SetConfigResponse{}
}

func (r *SysSetConfigRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type SysGetStatusRequest struct{}

func (r *SysGetStatusRequest) Method() string {
	return "Sys.GetStatus"
}

func (r *SysGetStatusRequest) Do(
	ctx context.Context,
	c mgrpc.MgRPC,
) (
	*SysStatus,
	*frame.Response,
	error,
) {
	resp := r.NewTypedResponse()
	raw, err := Do(ctx, c, r, resp)
	return resp, raw, err
}

func (r *SysGetStatusRequest) NewTypedResponse() *SysStatus {
	return &SysStatus{}
}

func (r *SysGetStatusRequest) NewResponse() any {
	return r.NewTypedResponse()
}

type SysConfig struct {
	// Device contains information about the device.
	Device *SysDeviceConfig `json:"device,omitempty"`

	// Location contains information about the current location of the device.
	Location *SysLocationConfig `json:"location,omitempty"`

	// Debug contains information about the device's debug logs.
	Debug *SysDebugConfig `json:"debug,omitempty"`

	// UI_Data contains user interface data. NOTE: the existance of this field is documented,
	// but not the contents.
	UI_Data json.RawMessage `json:"ui_data,omitempty"`

	// RPC_UDP is the configuration for the RPC over UDP.
	RPC_UDP *SysRPC_UDP_Config `json:"rpc_udp,omitempty"`

	// SNTP is configuration for the SNTP (time) server.
	SNTP *SysSNTP_Config `json:"sntp,omitempty"`

	// CfgRev is the configuration revision number. This number will be incremented for every
	// configuration change of a device component. If the new config value is the same as the
	// old one there will be no change of this property. Can not be modified explicitly by a call
	// to Sys.SetConfig.
	CfgRev int `json:"cfg_rev,omitempty"`
}

type SysDeviceConfig struct {
	// Name of the device.
	Name *string `json:"name"`

	// EcoMode (experimental) decreases power consumption when set to true, at the cost of reduced
	// execution speed and increased network latency.
	EcoMode bool `json:"eco_mode"`

	// MAC base MAC address of the device (Read-only).
	Mac string `json:"mac"`

	// FW_ID is the build identifier of the current firmware image.
	FW_ID string `json:"fw_id"`

	// Profile is the name of the device profile (only applicable for multi-profile devices)
	Profile *string `json:"profile,omitempty"`

	// Discoverable if true, device is shown in 'Discovered devices'. If false, the device is hidden.
	Discoverable bool `json:"discoverable"`

	// AddOnType enables/disables addon board (if supported). Range of values: sensor, prooutput;
	// null to disable.
	AddOnType *string `json:"addon_type,omitempty"`
}

type SysLocationConfig struct {
	// TZ is the timezone or null if unavailable.
	TZ *string `json:"tz"`

	// Lat is the latitude in degress or null if unavailable.
	Lat *float64 `json:"lat"`

	// Lon is the longitude in degress or null if unavailable.
	Lon *float64 `json:"lon"`
}

type SysDebugConfig struct {
	// MQTT contains configuration of logs streamed over MQTT.
	MQTT SysDebugConfigMQTT `json:"mqtt"`

	// Websocket contains configuration of logs stream over websocket.
	Websocket SysDebugConfigWebsocket `json:"websocket"`

	// UDP contains configuration of logs streamed over UDP
	UDP SysDebugConfigUDP `json:"udp"`

	// Level is not documented but appears in output.
	Level int `json:"level"`

	// FileLevel is not documented but appears in output.
	FileLevel interface{} `json:"file_level"`
}

type SysDebugConfigMQTT struct {
	// Enable will cause logs to be exported over MQTT if true.
	Enable bool `json:"enable"`
}

type SysDebugConfigWebsocket struct {
	// Enable will cause logs to be exported over websocket if true.
	Enable bool `json:"enable"`
}

type SysDebugConfigUDP struct {
	// Addr is the address that the device log is streamed to (null to disable logs).
	Addr *string `json:"addr"`
}

type SysRPC_UDP_Config struct {
	// DstAddr is the destination address for UDP.
	DstAddr *string `json:"dst_addr"`

	// ListenPort is the port number for inbound UDP RPC channel, null disables. Restart is
	// required for changes to apply
	ListenPort *int `json:"listen_port"`
}

type SysSNTP_Config struct {
	// Server is the name of the SNTP (time) server.
	Server string `json:"server"`
}

type SysStatus struct {
	// Mac address of the device.
	Mac string `json:"mac"`

	// RestartRequired will be true if a restart or update requires a system restart before going
	// into effect.
	RestartRequired bool `json:"restart_required"`

	// Time in the format HH:MM (24-hour time format in the current timezone with leading zero). null
	// when time is not synced from NTP server.
	Time *string `json:"time"`

	// UnixTime is the timestamp (in UTC), null when time is not synced from NTP server.
	UnixTime *int `json:"unixtime"`

	// Uptime is the time in seconds since last reboot.
	Uptime float64 `json:"uptime"`

	// RamSize is the total size of the RAM in the system in Bytes.
	RamSize int `json:"ram_size"`

	// RamFree is the size of the free RAM in the system in Bytes
	RamFree int `json:"ram_free"`

	// FS_Size is the total size of the file system in Bytes.
	FS_Size int `json:"fs_size"`

	// FS_Free is the size of the free file system in Bytes.
	FS_Free int `json:"fs_free"`

	// CfgRev is the configuration revision number.
	CfgRev int `json:"cfg_rev"`

	// KVRev is the KVS (Key-Value Store) revision number.
	KVRev int `json:"kvs_rev"`

	// SchedulesRev is the revision number, present if schedules are enabled.
	ScheduleRev *int `json:"schedule_rev"`

	// WebhookRev is the revision number, present if webhooks are enabled.
	WebhookRev *int `json:"webhook_rev"`

	// Information about available updates, similar to the one returned by Shelly.CheckForUpdate
	// (empty object: {}, if no updates available). This information is automatically updated every
	// 24 hours. Note that build_id and url for an update are not displayed here
	AvailableUpdates *AvailableUpdates `json:"available_updates,omitempty"`

	// WakeUpReason contains information about boot type and cause (only for battery-operated devices).
	WakeUpReason *WakeUpReason `json:"wakeup_reason,omitempty"`

	// Period (in seconds) at which device wakes up and sends "keep-alive" packet to cloud,
	// readonly. Count starts from last full wakeup.
	WakeUpPeriod float64 `json:"wakeup_period,omitempty"`

	// SafeMode is True if device is oprating in Safe Mode and is only present in this mode.
	SafeMode *bool `json:"safe_mode,omitempty"`

	// ResetReason is not documented, but appears in 1.1.0 responses.
	ResetReason *int `json:"reset_reason,omitempty"`
}

type WakeUpReason struct {
	// Boot type, one of: poweron, software_restart, deepsleep_wake, internal (e.g. brownout
	// detection, watchdog timeout, etc.), unknown.
	Boot string `json:"boot"`

	// Boot cause, one of: button, usb, periodic, status_update, alarm, alarm_test, undefined
	// (in case of deep sleep, reset was not caused by exit from deep sleep).
	Cause string `json:"cause"`
}

type AvailableUpdates struct {
	// Stable indicates the new stable version of the firmware.
	Stable *FirmwareUpdateVersion `json:"stable,omitempty"`

	// Beta indicates the new beta version of the firmware.
	Beta *FirmwareUpdateVersion `json:"beta,omitempty"`
}
