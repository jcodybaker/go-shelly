package shelly

// NotifyStatus implements the NotifyStatus and NotifyFullStatus payload types.
type NotifyStatus struct {
	ShellyGetStatusResponse

	// TS is the UTC unix timestamp when the status report was generated.
	TS float64 `json:"ts"`
}