package shelly

type SetConfigResponse struct {
	// RestartRequired is true if the system must be restarted for a pending change to take effect.
	RestartRequired bool `json:"restart_required,omitempty"`
}
