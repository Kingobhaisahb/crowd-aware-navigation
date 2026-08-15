package routing

type ZoneStatus string

const (
	Green  ZoneStatus = "green"
	Yellow ZoneStatus = "yellow"
	Red    ZoneStatus = "red"
)

type Point struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CrowdZone struct {
	ID     string     `json:"id"`
	Center Point      `json:"center"`
	Status ZoneStatus `json:"status"`
}

type RouteDecision struct {
	Route  []Point `json:"route"`
	Reason string  `json:"reason"`
}