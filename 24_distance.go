package main

import (
	"fmt"
	"math"
)

// Point
type Point struct {
	x float64
	y float64
}

// NewPoint
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

// Distance calculates the distance between two points
func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow((p2.X()-p1.X()), 2) + math.Pow((p2.Y()-p1.Y()), 2))
}

func main() {
	p1 := NewPoint(-7.12, 2)
	p2 := NewPoint(7.12, 4)

	fmt.Printf("Distance between two points: %.2f\n", Distance(p1, p2))
}
