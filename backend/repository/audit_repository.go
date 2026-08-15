package repository

import (
	"database/sql"
	"encoding/json"

	"crowd-aware-navigation/backend/internal/audit"
)

type AuditRepository struct {
	db *sql.DB
}

func NewAuditRepository(db *sql.DB) *AuditRepository {
	return &AuditRepository{
		db: db,
	}
}

func (r *AuditRepository) Save(record audit.AuditRecord) error {
	crowdZones, err := json.Marshal(record.CrowdZones)
	if err != nil {
		return err
	}

	route, err := json.Marshal(record.RecommendedRoute)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`
		INSERT INTO audit_logs (
			id,
			user_id,
			event_timestamp,
			latitude,
			longitude,
			crowd_zones,
			reason,
			recommended_route
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE id = id
	`,
		record.ID,
		record.UserID,
		record.EventTimestamp,
		record.UserLocation.Latitude,
		record.UserLocation.Longitude,
		string(crowdZones),
		record.Reason,
		string(route),
	)

	return err
}