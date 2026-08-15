package events

import (
	"sort"

	"crowd-aware-navigation/backend/internal/state"
)

func Replay(events []LocationEvent) state.UserState {
	if len(events) == 0 {
		return state.UserState{}
	}

	// Remove duplicate EventIDs first.
	unique := make(map[string]LocationEvent)

	for _, event := range events {
		if _, exists := unique[event.EventID]; exists {
			continue
		}

		unique[event.EventID] = event
	}

	ordered := make([]LocationEvent, 0, len(unique))

	for _, event := range unique {
		ordered = append(ordered, event)
	}

	// Deterministic ordering.
	sort.Slice(ordered, func(i, j int) bool {
		if ordered[i].SequenceNumber != ordered[j].SequenceNumber {
			return ordered[i].SequenceNumber < ordered[j].SequenceNumber
		}

		if !ordered[i].Timestamp.Equal(ordered[j].Timestamp) {
			return ordered[i].Timestamp.Before(ordered[j].Timestamp)
		}

		return ordered[i].EventID < ordered[j].EventID
	})

	// The last event in deterministic order represents
	// the reconstructed current state.
	final := ordered[len(ordered)-1]

	return state.UserState{
		UserID:         final.UserID,
		Latitude:       final.Latitude,
		Longitude:      final.Longitude,
		DeviceID:       final.DeviceID,
		Timestamp:      final.Timestamp,
		SequenceNumber: final.SequenceNumber,
		EventID:        final.EventID,
	}
}