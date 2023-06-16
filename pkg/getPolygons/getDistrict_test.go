package getPolygons

import (
	"log"
	"surge/config"
	"testing"
)

func TestReturnPolygons(t *testing.T) {
	err := config.InitCnf("../../test-config.json")
	tehranList, err := ReturnPolygons()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	log.Println(tehranList)
	if len(tehranList.Districts) != 22 {
		t.Errorf("Expected 22 but got %d", len(tehranList.Districts))
	}
	if len(tehranList.Polygons) != 22 {
		t.Errorf("Expected 22 but got %d", len(tehranList.Polygons))
	}
}
