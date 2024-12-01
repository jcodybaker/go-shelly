package shelly

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShellyGetStatusResponseUnmarshall(t *testing.T) {
	tcs := []struct {
		name   string
		input  string
		expect ShellyGetStatusResponse
	}{
		{
			name: "pro pm 4",
			input: `{
				"ble": {},
				"cloud": {
				  "connected": true
				},
				"eth": {
				  "ip": null
				},
				"input:0": {
				  "id": 0,
				  "state": false
				},
				"input:1": {
				  "id": 1,
				  "state": false
				},
				"input:2": {
				  "id": 2,
				  "state": false
				},
				"input:3": {
				  "id": 3,
				  "state": false
				},
				"mqtt": {
				  "connected": false
				},
				"switch:0": {
				  "id": 0,
				  "source": "timer",
				  "output": false,
				  "apower": 0,
				  "voltage": 120.8,
				  "freq": 60,
				  "current": 0,
				  "pf": 0,
				  "aenergy": {
					"total": 1342.238,
					"by_minute": [
					  0,
					  0,
					  0
					],
					"minute_ts": 1703811193
				  },
				  "ret_aenergy": {
					"total": 0,
					"by_minute": [
					  0,
					  0,
					  0
					],
					"minute_ts": 1703811193
				  },
				  "temperature": {
					"tC": 41.3,
					"tF": 106.3
				  }
				},
				"switch:1": {
				  "id": 1,
				  "source": "HTTP_in",
				  "output": true,
				  "apower": 83.9,
				  "voltage": 120.8,
				  "freq": 60,
				  "current": 1.143,
				  "pf": 0.61,
				  "aenergy": {
					"total": 102650.773,
					"by_minute": [
					  344.204,
					  1475.177,
					  1474.888
					],
					"minute_ts": 1703811193
				  },
				  "ret_aenergy": {
					"total": 0,
					"by_minute": [
					  0,
					  0,
					  0
					],
					"minute_ts": 1703811193
				  },
				  "temperature": {
					"tC": 41.3,
					"tF": 106.3
				  }
				},
				"switch:2": {
				  "id": 2,
				  "source": "HTTP_in",
				  "output": true,
				  "apower": 210.3,
				  "voltage": 120.9,
				  "freq": 60,
				  "current": 1.741,
				  "pf": 1,
				  "aenergy": {
					"total": 69346.948,
					"by_minute": [
					  840.825,
					  3605.178,
					  3624.834
					],
					"minute_ts": 1703811193
				  },
				  "ret_aenergy": {
					"total": 0,
					"by_minute": [
					  0,
					  0,
					  0
					],
					"minute_ts": 1703811193
				  },
				  "temperature": {
					"tC": 41.3,
					"tF": 106.3
				  }
				},
				"switch:3": {
				  "id": 3,
				  "source": "init",
				  "output": false,
				  "apower": 0,
				  "voltage": 120.9,
				  "freq": 60,
				  "current": 0,
				  "pf": 0,
				  "aenergy": {
					"total": 13.264,
					"by_minute": [
					  0,
					  0,
					  0
					],
					"minute_ts": 1703811193
				  },
				  "ret_aenergy": {
					"total": 0,
					"by_minute": [
					  0,
					  0,
					  0
					],
					"minute_ts": 1703811193
				  },
				  "temperature": {
					"tC": 41.3,
					"tF": 106.3
				  }
				},
				"sys": {
				  "mac": "C8F09E87D088",
				  "restart_required": false,
				  "time": "19:53",
				  "unixtime": 1703811195,
				  "uptime": 97431,
				  "ram_size": 241028,
				  "ram_free": 100452,
				  "fs_size": 524288,
				  "fs_free": 196608,
				  "cfg_rev": 26,
				  "kvs_rev": 1,
				  "schedule_rev": 0,
				  "webhook_rev": 0,
				  "available_updates": {},
				  "reset_reason": 3
				},
				"ui": {},
				"wifi": {
				  "sta_ip": "192.168.1.24",
				  "status": "got ip",
				  "ssid": "PickleTown",
				  "rssi": -36
				},
				"ws": {
				  "connected": false
				}
			  }`,
			expect: ShellyGetStatusResponse{
				BLE: &BLEStatus{},
				System: &SysStatus{
					Mac:              "C8F09E87D088",
					RestartRequired:  false,
					Time:             StrPtr("19:53"),
					UnixTime:         IntPtr(1703811195),
					Uptime:           97431,
					RamSize:          241028,
					RamFree:          100452,
					FS_Size:          524288,
					FS_Free:          196608,
					CfgRev:           26,
					KVRev:            1,
					ScheduleRev:      IntPtr(0),
					WebhookRev:       IntPtr(0),
					AvailableUpdates: &AvailableUpdates{},
					ResetReason:      IntPtr(3),
				},
				Cloud: &CloudStatus{
					Connected: true,
				},
				MQTT: &MQTTStatus{
					Connected: false,
				},
				Wifi: &WifiStatus{
					StaIP:  StrPtr("192.168.1.24"),
					Status: "got ip",
					SSID:   StrPtr("PickleTown"),
					RRSI:   Float64Ptr(-36),
				},
				Ethernet: &EthStatus{
					IP: nil,
				},
				Inputs: []*InputStatus{
					{
						ID:    0,
						State: BoolPtr(false),
					},
					{
						ID:    1,
						State: BoolPtr(false),
					},
					{
						ID:    2,
						State: BoolPtr(false),
					},
					{
						ID:    3,
						State: BoolPtr(false),
					},
				},
				Switches: []*SwitchStatus{
					{
						ID:      0,
						Source:  StrPtr("timer"),
						Output:  BoolPtr(false),
						APower:  Float64Ptr(0),
						Voltage: Float64Ptr(120.8),
						Freq:    Float64Ptr(60),
						Current: Float64Ptr(0),
						PF:      Float64Ptr(0),
						AEnergy: &EnergyCounters{
							Total:    1342.238,
							ByMinute: []float64{0, 0, 0},
							MinuteTS: 1703811193,
						},
						RetAEnergy: &EnergyCounters{
							Total:    0,
							ByMinute: []float64{0, 0, 0},
							MinuteTS: 1703811193,
						},
						Temperature: &Temperature{
							C: Float64Ptr(41.3),
							F: Float64Ptr(106.3),
						},
					},
					{
						ID:      1,
						Source:  StrPtr("HTTP_in"),
						Output:  BoolPtr(true),
						APower:  Float64Ptr(83.9),
						Voltage: Float64Ptr(120.8),
						Freq:    Float64Ptr(60),
						Current: Float64Ptr(1.143),
						PF:      Float64Ptr(0.61),
						AEnergy: &EnergyCounters{
							Total: 102650.773,
							ByMinute: []float64{
								344.204,
								1475.177,
								1474.888,
							},
							MinuteTS: 1703811193,
						},
						RetAEnergy: &EnergyCounters{
							Total:    0,
							ByMinute: []float64{0, 0, 0},
							MinuteTS: 1703811193,
						},
						Temperature: &Temperature{
							C: Float64Ptr(41.3),
							F: Float64Ptr(106.3),
						},
					},
					{
						ID:      2,
						Source:  StrPtr("HTTP_in"),
						Output:  BoolPtr(true),
						APower:  Float64Ptr(210.3),
						Voltage: Float64Ptr(120.9),
						Freq:    Float64Ptr(60),
						Current: Float64Ptr(1.741),
						PF:      Float64Ptr(1),
						AEnergy: &EnergyCounters{
							Total: 69346.948,
							ByMinute: []float64{
								840.825,
								3605.178,
								3624.834,
							},
							MinuteTS: 1703811193,
						},
						RetAEnergy: &EnergyCounters{
							Total:    0,
							ByMinute: []float64{0, 0, 0},
							MinuteTS: 1703811193,
						},
						Temperature: &Temperature{
							C: Float64Ptr(41.3),
							F: Float64Ptr(106.3),
						},
					},
					{
						ID:      3,
						Source:  StrPtr("init"),
						Output:  BoolPtr(false),
						APower:  Float64Ptr(0),
						Voltage: Float64Ptr(120.9),
						Freq:    Float64Ptr(60),
						Current: Float64Ptr(0),
						PF:      Float64Ptr(0),
						AEnergy: &EnergyCounters{
							Total: 13.264,
							ByMinute: []float64{
								0, 0, 0,
							},
							MinuteTS: 1703811193,
						},
						RetAEnergy: &EnergyCounters{
							Total:    0,
							ByMinute: []float64{0, 0, 0},
							MinuteTS: 1703811193,
						},
						Temperature: &Temperature{
							C: Float64Ptr(41.3),
							F: Float64Ptr(106.3),
						},
					},
				},
			},
		},
		{
			name: "pro 3",
			input: `{
				"ble": {},
				"cloud": {
				  "connected": true
				},
				"eth": {
				  "ip": null
				},
				"input:0": {
				  "id": 0,
				  "state": false
				},
				"input:1": {
				  "id": 1,
				  "state": false
				},
				"input:2": {
				  "id": 2,
				  "state": false
				},
				"mqtt": {
				  "connected": false
				},
				"switch:0": {
				  "id": 0,
				  "source": "init",
				  "output": false,
				  "temperature": {
					"tC": 35.7,
					"tF": 96.2
				  }
				},
				"switch:1": {
				  "id": 1,
				  "source": "timer",
				  "output": false,
				  "temperature": {
					"tC": 35.7,
					"tF": 96.2
				  }
				},
				"switch:2": {
				  "id": 2,
				  "source": "timer",
				  "output": false,
				  "temperature": {
					"tC": 35.7,
					"tF": 96.2
				  }
				},
				"sys": {
				  "mac": "C8F09E883630",
				  "restart_required": false,
				  "time": "19:52",
				  "unixtime": 1703811156,
				  "uptime": 98059,
				  "ram_size": 243420,
				  "ram_free": 104384,
				  "fs_size": 524288,
				  "fs_free": 212992,
				  "cfg_rev": 16,
				  "kvs_rev": 0,
				  "schedule_rev": 0,
				  "webhook_rev": 0,
				  "available_updates": {},
				  "reset_reason": 3
				},
				"wifi": {
				  "sta_ip": "192.168.1.23",
				  "status": "got ip",
				  "ssid": "PickleTown",
				  "rssi": -22
				},
				"ws": {
				  "connected": false
				}
			  }`,
			expect: ShellyGetStatusResponse{
				BLE: &BLEStatus{},
				System: &SysStatus{
					Mac:              "C8F09E883630",
					RestartRequired:  false,
					Time:             StrPtr("19:52"),
					UnixTime:         IntPtr(1703811156),
					Uptime:           98059,
					RamSize:          243420,
					RamFree:          104384,
					FS_Size:          524288,
					FS_Free:          212992,
					CfgRev:           16,
					KVRev:            0,
					ScheduleRev:      IntPtr(0),
					WebhookRev:       IntPtr(0),
					AvailableUpdates: &AvailableUpdates{},
					ResetReason:      IntPtr(3),
				},
				Cloud: &CloudStatus{
					Connected: true,
				},
				MQTT: &MQTTStatus{
					Connected: false,
				},
				Wifi: &WifiStatus{
					StaIP:  StrPtr("192.168.1.23"),
					Status: "got ip",
					SSID:   StrPtr("PickleTown"),
					RRSI:   Float64Ptr(-22),
				},
				Ethernet: &EthStatus{
					IP: nil,
				},
				Inputs: []*InputStatus{
					{
						ID:    0,
						State: BoolPtr(false),
					},
					{
						ID:    1,
						State: BoolPtr(false),
					},
					{
						ID:    2,
						State: BoolPtr(false),
					},
				},
				Switches: []*SwitchStatus{
					{
						ID:     0,
						Source: StrPtr("init"),
						Output: BoolPtr(false),
						Temperature: &Temperature{
							C: Float64Ptr(35.7),
							F: Float64Ptr(96.2),
						},
					},
					{
						ID:     1,
						Source: StrPtr("timer"),
						Output: BoolPtr(false),
						Temperature: &Temperature{
							C: Float64Ptr(35.7),
							F: Float64Ptr(96.2),
						},
					},
					{
						ID:     2,
						Source: StrPtr("timer"),
						Output: BoolPtr(false),
						Temperature: &Temperature{
							C: Float64Ptr(35.7),
							F: Float64Ptr(96.2),
						},
					},
				},
			},
		},
		{
			name: "shelly plus ht",
			input: `{
				"ble": {},
				"cloud": {
					"connected": true
				},
				"devicepower:0": {
					"id": 0,
					"battery": {
						"V": 0.43,
						"percent": 0
					},
					"external": {
						"present": true
					}
				},
				"ht_ui": {},
				"humidity:0": {
					"id": 0,
					"rh": 59.4
				},
				"mqtt": {
					"connected": true
				},
				"sys": {
					"mac": "C049EF8BB8F8",
					"restart_required": false,
					"time": "09:29",
					"unixtime": 1733063380,
					"uptime": 6,
					"ram_size": 247452,
					"ram_free": 159420,
					"fs_size": 458752,
					"fs_free": 176128,
					"cfg_rev": 15,
					"kvs_rev": 0,
					"webhook_rev": 2,
					"available_updates": {},
					"wakeup_reason": {
						"boot": "deepsleep_wake",
						"cause": "status_update"
					},
					"wakeup_period": 600,
					"reset_reason": 8
				},
				"temperature:0": {
					"id": 0,
					"tC": 2.7,
					"tF": 36.8
				},
				"wifi": {
					"sta_ip": "192.168.1.199",
					"status": "got ip",
					"ssid": "PickleTown_Garage",
					"rssi": -35
				},
				"ws": {
					"connected": false
				}
			}`,
			expect: ShellyGetStatusResponse{
				BLE: &BLEStatus{},
				System: &SysStatus{
					Mac:              "C049EF8BB8F8",
					RestartRequired:  false,
					Time:             StrPtr("09:29"),
					UnixTime:         IntPtr(1733063380),
					Uptime:           6,
					RamSize:          247452,
					RamFree:          159420,
					FS_Size:          458752,
					FS_Free:          176128,
					CfgRev:           15,
					KVRev:            0,
					WebhookRev:       IntPtr(2),
					AvailableUpdates: &AvailableUpdates{},
					ResetReason:      IntPtr(8),
					WakeUpReason: &WakeUpReason{
						Boot:  "deepsleep_wake",
						Cause: "status_update",
					},
					WakeUpPeriod: 600,
				},
				DevicePowers: []*DevicePowerStatus{
					{
						ID: 0,
						Battery: &DevicePowerBatteryStatus{
							V:       Float64Ptr(0.43),
							Percent: Float64Ptr(0),
						},
						External: &DevicePowerExternalStatus{
							Present: true,
						},
					},
				},
				Humidities: []*HumidityStatus{
					{
						ID: 0,
						RH: Float64Ptr(59.4),
					},
				},
				Temperatures: []*TemperatureStatus{
					{
						ID: 0,
						TC: Float64Ptr(2.7),
						TF: Float64Ptr(36.8),
					},
				},
				Cloud: &CloudStatus{
					Connected: true,
				},
				MQTT: &MQTTStatus{
					Connected: true,
				},
				Wifi: &WifiStatus{
					StaIP:  StrPtr("192.168.1.199"),
					Status: "got ip",
					SSID:   StrPtr("PickleTown_Garage"),
					RRSI:   Float64Ptr(-35),
				},
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			var got ShellyGetStatusResponse
			require.NoError(t, json.Unmarshal([]byte(tc.input), &got))
			assert.Equal(t, tc.expect, got)
		})
	}
}
