package main

import "math"

func add(a, b int) int {
	return a+b
}

func subtract(a, b int) int {
	return a-b
}

func multiply(a, b int) int {
	return a*b
}

func divide(a, b int) int {
	return a/b
}

func sin(a float64) float64 {
	sinVal := math.Sin(a)
	return math.Round(sinVal*100)/100
}

func cos(a float64) float64 {
	cosVal := math.Cos(a)
	return math.Round(cosVal*100)/100
}

func squareRoot(a float64) float64 {
	return math.Sqrt(a)
}