package recomputation

import (
	"time"

	"crowd-aware-navigation/backend/internal/audit"
	"crowd-aware-navigation/backend/internal/routing"
)

type Service struct {
	auditor *audit.Recorder
}

func NewService(auditor *audit.Recorder) *Service {
	return &Service{
		auditor: auditor,
	}
}

type Request struct {
	UserID          string
	EventTimestamp  time.Time
	UserLocation    routing.Point
	Destination     routing.Point
	CrowdZones      []routing.CrowdZone
	Reason          string
}

type Result struct {
	Decision routing.RouteDecision
	Audit    audit.AuditRecord
}

func (s *Service) Recompute(request Request) Result {
	decision := routing.ComputeRoute(
		request.UserLocation,
		request.Destination,
		request.CrowdZones,
	)

	if request.Reason != "" {
		decision.Reason = request.Reason
	}

	record := audit.NewRecord(
		request.UserID,
		request.EventTimestamp,
		request.UserLocation,
		request.CrowdZones,
		decision,
	)

	s.auditor.Record(record)

	return Result{
		Decision: decision,
		Audit:    record,
	}
}