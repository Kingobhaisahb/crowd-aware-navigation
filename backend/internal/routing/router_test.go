package routing

import "testing"

func TestComputeRouteAvoidsRedZone(t *testing.T) {
	start := Point{
		Latitude:  28.6467,
		Longitude: 77.3452,
	}

	destination := Point{
		Latitude: 28.6500,
		Longitude: 77.3500,
	}

	zones := []CrowdZone{
		{
			ID: "zone-1",
			Center: Point{
				Latitude:  28.6468,
				Longitude: 77.3453,
			},
			Status: Red,
		},
	}

	result := ComputeRoute(start, destination, zones)

	if len(result.Route) != 3 {
		t.Fatalf(
			"expected alternative route with 3 points, got %d",
			len(result.Route),
		)
	}

	if result.Reason == "" {
		t.Fatal("expected route recomputation reason")
	}
}

func TestComputeRouteDirectWhenSafe(t *testing.T) {
	start := Point{
		Latitude:  28.6467,
		Longitude: 77.3452,
	}

	destination := Point{
		Latitude: 28.6500,
		Longitude: 77.3500,
	}

	result := ComputeRoute(
		start,
		destination,
		[]CrowdZone{},
	)

	if len(result.Route) != 2 {
		t.Fatalf(
			"expected direct route with 2 points, got %d",
			len(result.Route),
		)
	}

	if result.Reason != "Direct route is safe" {
		t.Fatalf("unexpected reason: %s", result.Reason)
	}
}