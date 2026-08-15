package audit

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"crowd-aware-navigation/backend/internal/routing"
)

type AuditRecord struct {
	ID              string              `json:"id"`
	UserID          string              `json:"user_id"`
	EventTimestamp  time.Time           `json:"event_timestamp"`
	UserLocation    routing.Point       `json:"user_location"`
	CrowdZones      []routing.CrowdZone  `json:"crowd_zones"`
	Reason          string              `json:"reason"`
	RecommendedRoute []routing.Point     `json:"recommended_route"`
}

func NewRecord(
	userID string,
	eventTimestamp time.Time,
	userLocation routing.Point,
	crowdZones []routing.CrowdZone,
	decision routing.RouteDecision,
) AuditRecord {

	record := AuditRecord{
		UserID:           userID,
		EventTimestamp:   eventTimestamp,
		UserLocation:     userLocation,
		CrowdZones:       crowdZones,
		Reason:           decision.Reason,
		RecommendedRoute: decision.Route,
	}

	record.ID = generateID(record)

	return record
}

func generateID(record AuditRecord) string {
	record.ID = ""

	data, _ := json.Marshal(record)

	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:])
}