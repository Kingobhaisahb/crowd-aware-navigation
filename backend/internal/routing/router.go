package routing

func ComputeRoute(
	start Point,
	destination Point,
	zones []CrowdZone,
) RouteDecision {

	for _, zone := range zones {
		if zone.Status == Red && isNear(start, zone.Center) {
			return RouteDecision{
				Route: []Point{
					start,
					{
						Latitude:  start.Latitude,
						Longitude: destination.Longitude,
					},
					destination,
				},
				Reason: "Start location is inside or near a Red crowd zone",
			}
		}
	}

	return RouteDecision{
		Route: []Point{
			start,
			destination,
		},
		Reason: "Direct route is safe",
	}
}

func isNear(a Point, b Point) bool {
	const threshold = 0.01

	latDiff := a.Latitude - b.Latitude
	if latDiff < 0 {
		latDiff = -latDiff
	}

	lonDiff := a.Longitude - b.Longitude
	if lonDiff < 0 {
		lonDiff = -lonDiff
	}

	return latDiff <= threshold && lonDiff <= threshold
}
