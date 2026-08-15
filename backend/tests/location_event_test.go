package tests

import (
	"testing"
	"time"

	"crowd-aware-navigation/backend/models"
)

func TestLocationEventCreation(t *testing.T) {
	event := models.LocationEvent{
		EventID:        "event-001",
		UserID:         "user-001",
		DeviceID:       "device-001",
		Latitude:       28.6139,
		Longitude:      77.2090,
		Timestamp:      time.Unix(1000, 0),
		SequenceNumber: 1,
	}

	if event.EventID == "" {
		t.Fatal("event ID should not be empty")
	}

	if event.UserID == "" {
		t.Fatal("user ID should not be empty")
	}

	if event.DeviceID == "" {
		t.Fatal("device ID should not be empty")
	}

	if event.Latitude == 0 || event.Longitude == 0 {
		t.Fatal("location should be populated")
	}

	if event.SequenceNumber != 1 {
		t.Fatal("sequence number should be 1")
	}
}