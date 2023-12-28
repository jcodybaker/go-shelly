package shelly

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

// MDNSAppToDeviceSpecs translates the "app" field from an mDNS response to a device. For multi-profile
// devices, the aspirational capacity is returned but the actual specs depend on the active profile.
func MDNSAppToDeviceSpecs(mdnsApp string) (DeviceSpecs, bool) {
	switch mdnsApp {
	// These device "app" values have been confirmed.
	case "Pro3":
		return DeviceSpecs{
			Inputs:             3,
			Switches:           3,
			Scripts:            10,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
		}, true
	case "Pro4PM":
		return DeviceSpecs{
			Inputs:             4,
			Switches:           4,
			SwitchEnergy:       true,
			Scripts:            10,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
			UI:                 true,
		}, true
	case "PlugUS", "PlugUK", "PlugS", "PlugIT":
		return DeviceSpecs{
			Inputs:             1,
			Switches:           1,
			Scripts:            10,
			SwitchEnergy:       true,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, true
	case "PlusHT":
		return DeviceSpecs{
			Humidity:              true,
			Temperature:           true,
			HumidityTemperatureUI: true,
			DevicePower:           true,
			Wifi:                  true,
			BluetoothLowEnergy:    true,
		}, true

	// These names are best guesses.
	case "Plus1", "Plus1Mini":
		return DeviceSpecs{
			Inputs:             1,
			Switches:           1,
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, true
	case "Plus1PM", "Plus1PMMini":
		return DeviceSpecs{
			Inputs:             1,
			Switches:           1,
			SwitchEnergy:       true,
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, true
	case "Plus2PM":
		return DeviceSpecs{
			Profiles:           []string{"switch", "cover"},
			Covers:             1,
			Inputs:             2,
			Switches:           2,
			SwitchEnergy:       true,
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, true
	case "PlusI4":
		return DeviceSpecs{
			Inputs:             4,
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, true
	case "PlusSmoke":
		return DeviceSpecs{
			Wifi:               true,
			BluetoothLowEnergy: true,
			DevicePower:        true,
			Smoke:              true,
		}, true
	case "PlusWallDimmer":
		return DeviceSpecs{
			Lights:             1,
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
			WallDimmerUI:       true,
		}, true
	case "Plus0-10VDimmer":
		return DeviceSpecs{
			Lights:             1,
			Inputs:             2,
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, true
	case "PlusPMMini":
		return DeviceSpecs{
			PM1:                true,
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, true
	case "Pro1":
		return DeviceSpecs{
			Inputs:             2,
			Switches:           1,
			Scripts:            10,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
		}, true
	case "Pro1PM":
		return DeviceSpecs{
			Inputs:             2,
			Switches:           1,
			SwitchEnergy:       true,
			Scripts:            10,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
		}, true
	case "Pro2":
		return DeviceSpecs{
			Inputs:             2,
			Switches:           2,
			Scripts:            10,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
		}, true
	case "Pro2PM":
		return DeviceSpecs{
			Profiles:           []string{"switch", "cover"},
			Covers:             1,
			Inputs:             2,
			Switches:           2,
			Scripts:            10,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
		}, true
	case "ProDualCoverPM":
		return DeviceSpecs{
			Covers:             2,
			Inputs:             4,
			Scripts:            10,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
			UI:                 true,
		}, true
	case "ProEM":
		return DeviceSpecs{
			Switches:           1,
			EM1:                2,
			EM1Data:            2,
			Scripts:            10,
			ModBus:             true,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
		}, true
	case "Pro3EM", "Pro3EM400":
		return DeviceSpecs{
			Switches:           1,
			EM:                 true,
			EMData:             true,
			Scripts:            10,
			ModBus:             true,
			Wifi:               true,
			Ethernet:           true,
			BluetoothLowEnergy: true,
		}, true
	case "BLEGateway":
		return DeviceSpecs{
			Scripts:            10,
			Wifi:               true,
			BluetoothLowEnergy: true,
		}, true
	}
	return DeviceSpecs{}, false
}
