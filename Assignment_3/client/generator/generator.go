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
