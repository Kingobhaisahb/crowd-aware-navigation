package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"crowd-aware-navigation/backend/internal/routing"
	"crowd-aware-navigation/backend/services/recomputation"
)

type RecomputationHandler struct {
	service *recomputation.Service
}

func NewRecomputationHandler(
	service *recomputation.Service,
) *RecomputationHandler {
	return &RecomputationHandler{
		service: service,
	}
}

type RecomputeRequest struct {
	UserID         string               `json:"user_id"`
	EventTimestamp time.Time            `json:"event_timestamp"`
	UserLocation   routing.Point        `json:"user_location"`
	Destination    routing.Point        `json:"destination"`
	CrowdZones     []routing.CrowdZone  `json:"crowd_zones"`
	Reason         string               `json:"reason"`
}

func (h *RecomputationHandler) Recompute(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method != http.MethodPost {
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var request RecomputeRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	if request.UserID == "" {
		http.Error(
			w,
			"user_id is required",
			http.StatusBadRequest,
		)
		return
	}

	result := h.service.Recompute(
		recomputation.Request{
			UserID:         request.UserID,
			EventTimestamp: request.EventTimestamp,
			UserLocation:   request.UserLocation,
			Destination:    request.Destination,
			CrowdZones:     request.CrowdZones,
			Reason:         request.Reason,
		},
	)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(
			w,
			"failed to encode response",
			http.StatusInternalServerError,
		)
		return
	}
}