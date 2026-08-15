package events

import "sort"

type EventStatus string

const (
	StatusAccepted  EventStatus = "accepted"
	StatusDuplicate EventStatus = "duplicate"
	StatusLate      EventStatus = "late"
	StatusConflict  EventStatus = "conflict"
)

type ProcessResult struct {
	Event  LocationEvent `json:"event"`
	Status EventStatus   `json:"status"`
	Reason string        `json:"reason"`
}

// ProcessEvents removes duplicates and deterministically orders
// events for state reconstruction.
func ProcessEvents(events []LocationEvent) []ProcessResult {
	if len(events) == 0 {
		return []ProcessResult{}
	}

	// Sort deterministically.
	ordered := make([]LocationEvent, len(events))
	copy(ordered, events)

	sort.Slice(ordered, func(i, j int) bool {
		if ordered[i].SequenceNumber != ordered[j].SequenceNumber {
			return ordered[i].SequenceNumber < ordered[j].SequenceNumber
		}

		if !ordered[i].Timestamp.Equal(ordered[j].Timestamp) {
			return ordered[i].Timestamp.Before(ordered[j].Timestamp)
		}

		return ordered[i].EventID < ordered[j].EventID
	})

	seen := make(map[string]bool)
	results := make([]ProcessResult, 0, len(ordered))

	var highestSequence int

	for _, event := range ordered {
		// Duplicate event.
		if seen[event.EventID] {
			results = append(results, ProcessResult{
				Event:  event,
				Status: StatusDuplicate,
				Reason: "event ID already processed",
			})
			continue
		}

		seen[event.EventID] = true

		// Same sequence but different event = conflict.
		conflict := false

		for _, previous := range results {
			if previous.Status == StatusDuplicate {
				continue
			}

			if previous.Event.SequenceNumber == event.SequenceNumber &&
				previous.Event.EventID != event.EventID {
				conflict = true
				break
			}
		}

		if conflict {
			results = append(results, ProcessResult{
				Event:  event,
				Status: StatusConflict,
				Reason: "multiple events share the same sequence number",
			})
			continue
		}

		// Sequence lower than the highest already accepted.
		if event.SequenceNumber < highestSequence {
			results = append(results, ProcessResult{
				Event:  event,
				Status: StatusLate,
				Reason: "event arrived with an older sequence number",
			})
			continue
		}

		highestSequence = event.SequenceNumber

		results = append(results, ProcessResult{
			Event:  event,
			Status: StatusAccepted,
			Reason: "event accepted for replay",
		})
	}

	return results
}