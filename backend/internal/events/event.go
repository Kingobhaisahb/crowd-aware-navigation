package events

import "time"

type LocationEvent struct {
	EventID        string    `json:"eventId"`
	UserID         string    `json:"userId"`
	DeviceID       string    `json:"deviceId"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	Timestamp      time.Time `json:"timestamp"`
	SequenceNumber int       `json:"sequenceNumber"`
}