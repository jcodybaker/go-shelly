package shelly

type CloudSetConfigRequest struct {
	Config CloudConfig `json:"config"`
}

func (r *CloudSetConfigRequest) Method() string {
	return "Cloud.SetConfig"
}

type CloudConfig struct {
	// Enable is true if cloud connection is enabled, false otherwise
	Enable bool `json:"enable"`

	// Server is the name of the server to which the device is connected (optional).
	Server *string `json:"server"`
}

type CloudGetConfigRequest struct{}

func (r *CloudGetConfigRequest) Method() string {
	return "Cloud.GetConfig"
}

type CloudStatus struct {
	Connected bool `json:"connected"`
}

type CloudGetStatusRequest struct{}

func (r *CloudGetStatusRequest) Method() string {
	return "Cloud.GetStatus"
}
