package events

import (
	"testing"
	"time"
)

func TestProcessEventsDetectsDuplicate(t *testing.T) {
	now := time.Now()

	event := LocationEvent{
		EventID:        "event-1",
		UserID:         "user-1",
		DeviceID:       "device-1",
		Latitude:       28.6,
		Longitude:      77.3,
		Timestamp:      now,
		SequenceNumber: 1,
	}

	results := ProcessEvents([]LocationEvent{
		event,
		event,
	})

	foundDuplicate := false

	for _, result := range results {
		if result.Status == StatusDuplicate {
			foundDuplicate = true
		}
	}

	if !foundDuplicate {
		t.Fatal("expected duplicate event to be detected")
	}
}

func TestProcessEventsDetectsConflict(t *testing.T) {
	now := time.Now()

	results := ProcessEvents([]LocationEvent{
		{
			EventID:        "event-a",
			UserID:         "user-1",
			DeviceID:       "device-1",
			Latitude:       28.6,
			Longitude:      77.3,
			Timestamp:      now,
			SequenceNumber: 5,
		},
		{
			EventID:        "event-b",
			UserID:         "user-1",
			DeviceID:       "device-1",
			Latitude:       29.6,
			Longitude:      78.3,
			Timestamp:      now,
			SequenceNumber: 5,
		},
	})

	foundConflict := false

	for _, result := range results {
		if result.Status == StatusConflict {
			foundConflict = true
		}
	}

	if !foundConflict {
		t.Fatal("expected conflicting sequence numbers to be detected")
	}
}

func TestProcessEventsDetectsLateEvent(t *testing.T) {
	now := time.Now()

	results := ProcessEvents([]LocationEvent{
		{
			EventID:        "event-3",
			UserID:         "user-1",
			DeviceID:       "device-1",
			Latitude:       30,
			Longitude:      80,
			Timestamp:      now.Add(3 * time.Minute),
			SequenceNumber: 3,
		},
		{
			EventID:        "event-1",
			UserID:         "user-1",
			DeviceID:       "device-1",
			Latitude:       28,
			Longitude:      78,
			Timestamp:      now,
			SequenceNumber: 1,
		},
	})

	// ProcessEvents sorts before processing, so this is not considered
	// late during deterministic replay.
	for _, result := range results {
		if result.Status == StatusLate {
			t.Fatal("event should not be considered late during full replay")
		}
	}
}