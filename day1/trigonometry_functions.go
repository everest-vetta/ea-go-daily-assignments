package main

import "math"

func sincos(first float64) (float64, float64) {
	sin, cos := math.Sincos(first * math.Pi / 180)
	return sin, cos
}

func tan(first float64) float64 {
	return math.Tan(first * math.Pi / 180)
}
