package main

import "math"

func percof(x, y float64) float64 {
	return (x / y) * 100
}

func sqroot(x float64) float64 {
	return math.Sqrt(x)
}

func deg2rad(x float64) float64 {
	return x / math.Pi * 180
}

func rad2deg(x float64) float64 {
	return x * 180 / math.Pi
}

func degsine(x float64) float64 {
	return math.Sin(deg2rad(x))
}

func radsine(x float64) float64 {
	return math.Sin(x)
}