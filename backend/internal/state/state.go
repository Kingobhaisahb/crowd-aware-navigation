package state

import "time"

type UserState struct {
	UserID         string    `json:"userId"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	DeviceID       string    `json:"deviceId"`
	Timestamp      time.Time `json:"timestamp"`
	SequenceNumber int       `json:"sequenceNumber"`
	EventID        string    `json:"eventId"`
}