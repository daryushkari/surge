package usecase

import (
	"surge/pkg/getPolygons"
	"testing"
)

type isPointTests struct {
	lat, lon float64
	expected bool
}

var tableTests = []isPointTests{
	{5, 5, true},
	{-5, -5, false},
	{0, 0, false},
}

func TestIsPointInPolygon(t *testing.T) {
	district := &getPolygons.DistrictPolygon{}
	district.Points = []*getPolygons.Point{
		{Latitude: 0, Longitude: 0},
		{Latitude: 0, Longitude: 10},
		{Latitude: 10, Longitude: -10},
		{Latitude: 20, Longitude: 0},
		{Latitude: 10, Longitude: 10},
	}

	for _, test := range tableTests {
		if output := isPointInPolygon(test.lon, test.lat, district); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}
