package recomputation

import (
	"testing"
	"time"

	"crowd-aware-navigation/backend/internal/audit"
	"crowd-aware-navigation/backend/internal/routing"
)

func TestRecomputeCreatesAuditRecord(t *testing.T) {
	auditor := audit.NewRecorder()
	service := NewService(auditor)

	result := service.Recompute(Request{
		UserID:         "user-1",
		EventTimestamp: time.Unix(1000, 0),
		UserLocation: routing.Point{
			Latitude:  28.6467,
			Longitude: 77.3452,
		},
		Destination: routing.Point{
			Latitude:  28.6500,
			Longitude: 77.3500,
		},
		CrowdZones: []routing.CrowdZone{
			{
				ID: "red-zone",
				Center: routing.Point{
					Latitude:  28.6480,
					Longitude: 77.3470,
				},
				Status: routing.Red,
			},
		},
		Reason: "user entered red zone",
	})

	if len(result.Decision.Route) == 0 {
		t.Fatal("expected a route to be generated")
	}

	if result.Audit.ID == "" {
		t.Fatal("expected an audit ID")
	}

	if auditor.Count() != 1 {
		t.Fatalf(
			"expected 1 audit record, got %d",
			auditor.Count(),
		)
	}
}

func TestRecomputeIsIdempotent(t *testing.T) {
	auditor := audit.NewRecorder()
	service := NewService(auditor)

	request := Request{
		UserID:         "user-1",
		EventTimestamp: time.Unix(1000, 0),
		UserLocation: routing.Point{
			Latitude:  28.6467,
			Longitude: 77.3452,
		},
		Destination: routing.Point{
			Latitude:  28.6500,
			Longitude: 77.3500,
		},
		CrowdZones: []routing.CrowdZone{
			{
				ID: "red-zone",
				Center: routing.Point{
					Latitude:  28.6480,
					Longitude: 77.3470,
				},
				Status: routing.Red,
			},
		},
		Reason: "crowd density changed",
	}

	first := service.Recompute(request)
	second := service.Recompute(request)

	if first.Audit.ID != second.Audit.ID {
		t.Fatal("expected identical inputs to produce identical audit IDs")
	}

	if auditor.Count() != 1 {
		t.Fatalf(
			"expected 1 audit record after duplicate recomputation, got %d",
			auditor.Count(),
		)
	}
}