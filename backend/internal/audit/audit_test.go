package audit

import (
	"testing"
	"time"

	"crowd-aware-navigation/backend/internal/routing"
)

func TestAuditRecordIsDeterministic(t *testing.T) {
	location := routing.Point{
		Latitude:  28.6467,
		Longitude: 77.3452,
	}

	zones := []routing.CrowdZone{
		{
			ID: "zone-1",
			Center: routing.Point{
				Latitude:  28.6468,
				Longitude: 77.3453,
			},
			Status: routing.Red,
		},
	}

	decision := routing.ComputeRoute(
		location,
		routing.Point{
			Latitude:  28.6500,
			Longitude: 77.3500,
		},
		zones,
	)

	timestamp := time.Unix(1000, 0)

	record1 := NewRecord(
		"user-1",
		timestamp,
		location,
		zones,
		decision,
	)

	record2 := NewRecord(
		"user-1",
		timestamp,
		location,
		zones,
		decision,
	)

	if record1.ID != record2.ID {
		t.Fatalf(
			"expected deterministic IDs, got %s and %s",
			record1.ID,
			record2.ID,
		)
	}
}

func TestRecorderIsIdempotent(t *testing.T) {
	recorder := NewRecorder()

	record := AuditRecord{
		ID: "test-record",
	}

	recorder.Record(record)
	recorder.Record(record)

	if recorder.Count() != 1 {
		t.Fatalf(
			"expected 1 record after duplicate insertion, got %d",
			recorder.Count(),
		)
	}
}