package main

import (
	"math"
)

const PI = 3.14156

func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * PI
}

func Rect(base, hight float64) float64 {
	return base * hight
}

func _TriangleEq(base, hight float64) float64 {
	return (base * hight) / 2
}
