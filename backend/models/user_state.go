package models

import "time"

type UserState struct {
	UserID          string    `json:"userId"`
	DeviceID        string    `json:"deviceId"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	LastTimestamp   time.Time `json:"lastTimestamp"`
	LastSequence    int64     `json:"lastSequence"`
	LastEventID     string    `json:"lastEventId"`
}