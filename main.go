package main

import (
	"fmt"
	"math"
)

type Color struct {
	red int
	green int
	blue int
}

type Point struct {
	y float64
	x float64
}

func find_slope(first *Point, second *Point) float64 {
	return (second.y - first.y) / (second.x - first.x)
}

func find_y(x float64, known_x []float64) float64 {
	_, fraction := math.Modf(x)
	if fraction == 0.0 {
		return known_x[int(x)]
	}

	left_x := int(math.Trunc(x))
	right_x := left_x + 1

	left_y := known_x[left_x]
	right_y := known_x[right_x]

	left_point := Point {float64(left_x), left_y}
	right_point := Point {float64(right_x), right_y}

	slope := find_slope(&left_point, &right_point)

	return left_y + (slope * (x - float64(left_x)))
}

func main() {
	fmt.Println("Hey")
}
