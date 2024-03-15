package generator

import (
	"math/rand"
	"time"
)
func GenerateWater() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(100)
}

func GenerateWind() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(100)
}

func DetermineStatus(water, wind int) string {
    var status string

    // Menentukan status berdasarkan nilai water
    switch {
    case water < 5:
        status = "Aman"
    case water >= 6 && water <= 8:
        status = "Siaga"
    default:
        status = "Bahaya"
    }

    // Menentukan status berdasarkan nilai wind
    switch {
    case wind < 6:
        status += " - Aman"
    case wind >= 7 && wind <= 15:
        status += " - Siaga"
    default:
        status += " - Bahaya"
    }

    return status
}