package main

import (
	"math"
	"testing"
)

func TestDistanceSameCoordinates(t *testing.T) {
	dist := distance(0, 0, 0, 0)
	if dist != 0 {
		t.Fatalf("Expected distance of 0 between identical coordinates")
	}
}

func TestDistanceInMiles(t *testing.T) {
	dist := distance(0, 0, 0.1, 0.1)
	if math.Round(dist) != 10 {
		t.Fatalf("Expected distance of 10 miles between (0, 0) and (0.1, 0.1)")
	}
}

func TestDistanceInKm(t *testing.T) {
	dist := distance(0, 0, 0.1, 0.1, "K")
	if math.Round(dist) != 16 {
		t.Fatalf("Expected distance of 16 km between (0, 0) and (0.1, 0.1)")
	}
}
