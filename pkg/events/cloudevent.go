package events

import "time"

type CloudEvent struct {
	ID          string      `json:"id"`
	Source      string      `json:"source"`
	SpecVersion string      `json:"specversion"`
	Type        string      `json:"type"`
	Time        time.Time   `json:"time"`
	DataContentType string  `json:"datacontenttype"`
	Data        interface{} `json:"data"`
}
