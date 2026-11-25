package shelly

import "testing"

func TestPlugMG3DeviceSpecs(t *testing.T) {
	spec, err := AppToDeviceSpecs("PlugMG3", "")
	if err != nil {
		t.Fatal(err)
	}
	if spec.Switches != 1 || !spec.SwitchEnergy {
		t.Fatalf("unexpected device specs: %+v", spec)
	}
}
