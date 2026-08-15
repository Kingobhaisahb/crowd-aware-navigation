package events

import (
	"testing"
	"time"
)

func TestReplayOrdersEventsBySequence(t *testing.T) {
	now := time.Now()

	events := []LocationEvent{
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
		{
			EventID:        "event-2",
			UserID:         "user-1",
			DeviceID:       "device-1",
			Latitude:       29,
			Longitude:      79,
			Timestamp:      now.Add(time.Minute),
			SequenceNumber: 2,
		},
	}

	result := Replay(events)

	if result.EventID != "event-3" {
		t.Fatalf(
			"expected event-3 as final event, got %s",
			result.EventID,
		)
	}

	if result.SequenceNumber != 3 {
		t.Fatalf(
			"expected sequence 3, got %d",
			result.SequenceNumber,
		)
	}
}

func TestReplayIgnoresDuplicateEvents(t *testing.T) {
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

	result := Replay([]LocationEvent{
		event,
		event,
	})

	if result.EventID != "event-1" {
		t.Fatalf("expected event-1, got %s", result.EventID)
	}
}

func TestReplayHandlesOutOfOrderEvents(t *testing.T) {
	now := time.Now()

	events := []LocationEvent{
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
		{
			EventID:        "event-2",
			UserID:         "user-1",
			DeviceID:       "device-1",
			Latitude:       29,
			Longitude:      79,
			Timestamp:      now.Add(time.Minute),
			SequenceNumber: 2,
		},
	}

	result := Replay(events)

	if result.SequenceNumber != 3 {
		t.Fatalf(
			"expected final sequence 3, got %d",
			result.SequenceNumber,
		)
	}
}

func TestReplayIsDeterministic(t *testing.T) {
	now := time.Now()

	events := []LocationEvent{
		{
			EventID:        "event-b",
			UserID:         "user-1",
			DeviceID:       "device-1",
			Latitude:       29,
			Longitude:      79,
			Timestamp:      now,
			SequenceNumber: 2,
		},
		{
			EventID:        "event-a",
			UserID:         "user-1",
			DeviceID:       "device-1",
			Latitude:       28,
			Longitude:      78,
			Timestamp:      now,
			SequenceNumber: 1,
		},
	}

	first := Replay(events)
	second := Replay(events)

	if first.EventID != second.EventID {
		t.Fatalf(
			"replay is not deterministic: %s != %s",
			first.EventID,
			second.EventID,
		)
	}

	if first.SequenceNumber != second.SequenceNumber {
		t.Fatalf(
			"replay is not deterministic: %d != %d",
			first.SequenceNumber,
			second.SequenceNumber,
		)
	}
}