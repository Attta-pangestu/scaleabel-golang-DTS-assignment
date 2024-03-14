package main

import (
	"math/rand"
	"time"
)

const (
	WaterBelowThreshold = "aman"
	WaterWarning        = "siaga"
	WaterDanger         = "bahaya"
	WindBelowThreshold  = "aman"
	WindWarning         = "siaga"
	WindDanger          = "bahaya"
)

func generateWater() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func generateWind() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func determineStatus(water, wind int) string {
	status := WaterBelowThreshold
	if water >= 6 && water <= 8 {
		status = WaterWarning
	} else if water > 8 {
		status = WaterDanger
	}

	if wind >= 7 && wind <= 15 {
		if status == WaterBelowThreshold {
			status = WindWarning
		} else if status == WaterWarning {
			status = WindDanger
		}
	} else if wind > 15 {
		status = WindDanger
	}

	return status
}