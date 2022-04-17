package events

import "time"

type EventBridgeEnvelope struct {
	Source     string      `json:"Source"`
	DetailType string      `json:"DetailType"`
	Detail     interface{} `json:"Detail"`
	EventBus   string      `json:"EventBusName"`
	Time       time.Time   `json:"Time"`
}
