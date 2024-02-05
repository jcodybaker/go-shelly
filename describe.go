package shelly

import "errors"

var (
	// ErrUnknownDeviceApp is returned when the app cannot be resolved to a spec.
	ErrUnknownDeviceApp = errors.New("unable to resolve device `app` to spec: unknown `app`")

	// ErrUnknownDeviceProfile is returned when a device
	ErrUnknownDeviceProfile = errors.New("unable to resolve device profile to spec: unknown profile")
)

// DeviceSpecs describes the device abilities.
type DeviceSpecs struct {
	// Profiles describes the possible profiles for the device.
	Profiles []string

	Inputs       int
	Switches     int
	SwitchEnergy bool
	Scripts      int
	Lights       int
	Covers       int

	Temperature bool
	Humidity    bool
	DevicePower bool
	Smoke       bool

	ModBus             bool
	Wifi               bool
	Ethernet           bool
	BluetoothLowEnergy bool

	UI                    bool
	WallDimmerUI          bool
	HumidityTemperatureUI bool

	// PM describes devices (Shelly Plus PM Mini) which have power-measurement
	// capabilities independent of a switch.
	PM1 bool
	// EM  (Energy Meter) component handles the data collection and processing
	// from energy meter devices like the ShellyPro3EM.
	EM     bool
	EMData bool
	// EM1 component handles the data collection and processing from energy meter
	// devices like the ShellyProEM.
	EM1     int
	EM1Data int
}

// IsMultiProfile will return true if this device supports multiple profiles. In practice
// this means the actual profile must be read to determine some of the devices capabilities,
// namely the number of switches and covers.
func (s *DeviceSpecs) IsMultiProfile() bool {
	return len(s.Profiles) > 0
}

// AppToDeviceSpecs translates the "app" field from an mDNS response to a device. The
// "profile" parameter is optional. For multi-profile devices, without a profile parameter
// the aspirational capacity is returned but the actual specs depend on the active profile.
func AppToDeviceSpecs(mdnsApp, profile string) (DeviceSpecs, error) {
	switch mdnsApp {
	// These device "app" values have been confirmed.
	case "Pro3":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Inputs:             3,
			Switches:           3,
			Scripts:            10,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
		}, nil
	case "Pro4PM":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Inputs:             4,
			Switches:           4,
			SwitchEnergy:       true,
			Scripts:            10,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
			UI:                 true,
		}, nil
	case "PlugUS", "PlugUK", "PlugS", "PlugIT":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Switches:           1,
			Scripts:            10,
			SwitchEnergy:       true,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, nil
	case "PlusHT":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Humidity:              true,
			Temperature:           true,
			HumidityTemperatureUI: true,
			DevicePower:           true,
			Wifi:                  true,
			BluetoothLowEnergy:    true,
		}, nil

	// These names are best guesses.
	case "Plus1", "Plus1Mini":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Inputs:             1,
			Switches:           1,
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, nil
	case "Plus1PM", "Plus1PMMini":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Inputs:             1,
			Switches:           1,
			SwitchEnergy:       true,
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, nil
	case "Plus2PM":
		switch profile {
		case "":
			return DeviceSpecs{
				Profiles:           []string{"switch", "cover"},
				Covers:             1,
				Inputs:             2,
				Switches:           2,
				SwitchEnergy:       true,
				Scripts:            10,
				Wifi:               true,
				BluetoothLowEnergy: true,
			}, nil
		case "switch":
			return DeviceSpecs{
				Profiles:           []string{"switch", "cover"},
				Inputs:             2,
				Switches:           2,
				SwitchEnergy:       true,
				Scripts:            10,
				Wifi:               true,
				BluetoothLowEnergy: true,
			}, nil
		case "cover":
			return DeviceSpecs{
				Profiles:           []string{"switch", "cover"},
				Inputs:             2,
				Covers:             1,
				SwitchEnergy:       true,
				Scripts:            10,
				Wifi:               true,
				BluetoothLowEnergy: true,
			}, nil
		default:
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
	case "PlusI4":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Inputs:             4,
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, nil
	case "PlusSmoke":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Wifi:               true,
			BluetoothLowEnergy: true,
			DevicePower:        true,
			Smoke:              true,
		}, nil
	case "PlusWallDimmer":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Lights:             1,
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
			WallDimmerUI:       true,
		}, nil
	case "Plus0-10VDimmer":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Lights:             1,
			Inputs:             2,
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, nil
	case "PlusPMMini":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			PM1:                true,
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, nil
	case "Pro1":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Inputs:             2,
			Switches:           1,
			Scripts:            10,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
		}, nil
	case "Pro1PM":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Inputs:             2,
			Switches:           1,
			SwitchEnergy:       true,
			Scripts:            10,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
		}, nil
	case "Pro2":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Inputs:             2,
			Switches:           2,
			Scripts:            10,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
		}, nil
	case "Pro2PM":
		switch profile {
		case "":
			return DeviceSpecs{
				Profiles:           []string{"switch", "cover"},
				Covers:             1,
				Inputs:             2,
				Switches:           2,
				Scripts:            10,
				Wifi:               true,
				Ethernet:           true,
				BluetoothLowEnergy: true,
			}, nil
		case "switch":
			return DeviceSpecs{
				Profiles:           []string{"switch", "cover"},
				Inputs:             2,
				Switches:           2,
				Scripts:            10,
				Wifi:               true,
				Ethernet:           true,
				BluetoothLowEnergy: true,
			}, nil
		case "cover":
			return DeviceSpecs{
				Profiles:           []string{"switch", "cover"},
				Covers:             1,
				Inputs:             2,
				Scripts:            10,
				Wifi:               true,
				Ethernet:           true,
				BluetoothLowEnergy: true,
			}, nil
		default:
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}

	case "ProDualCoverPM":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Covers:             2,
			Inputs:             4,
			Scripts:            10,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
			UI:                 true,
		}, nil
	case "ProEM":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Switches:           1,
			EM1:                2,
			EM1Data:            2,
			Scripts:            10,
			ModBus:             true,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
		}, nil
	case "Pro3EM", "Pro3EM400":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Switches:           1,
			EM:                 true,
			EMData:             true,
			Scripts:            10,
			ModBus:             true,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
		}, nil
	case "BLEGateway":
		if profile != "" {
			return DeviceSpecs{}, ErrUnknownDeviceProfile
		}
		return DeviceSpecs{
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, nil
	}
	return DeviceSpecs{}, ErrUnknownDeviceApp
}
