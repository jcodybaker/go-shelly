package shelly

import "encoding/json"

// NotifyStatus implements the NotifyStatus and NotifyFullStatus payload types.
type NotifyStatus struct {
	ShellyGetStatusResponse

	// TS is the UTC unix timestamp when the status report was generated.
	TS float64 `json:"ts"`
}

func (ns *NotifyStatus) Method() string {
	return "NotifyStatus"
}

func (r *NotifyStatus) UnmarshalJSON(b []byte) error {
	var onlyTS = struct {
		TS float64 `json:"ts"`
	}{}
	if err := json.Unmarshal(b, &onlyTS); err != nil {
		return err
	}
	r.TS = onlyTS.TS
	return r.ShellyGetStatusResponse.UnmarshalJSON(b)
}

type NotifyEvent struct {
	// TS is the UTC unix timestamp when the event report was generated.
	TS float64 `json:"ts"`

	Events []Event `json:"events"`
}

func (ne *NotifyEvent) Method() string {
	return "NotifyEvent"
}

type Event struct {
	// TS is the UTC unix timestamp when the status report was generated.
	TS float64 `json:"ts"`

	// Component key (component_type[:id], e.g. switch:0; wifi
	Component string `json:"component,omitempty"`

	// ID of the component instance.
	ID int `json:"id,omitempty"`

	// Event name.
	Event string `json:"event"`
}
