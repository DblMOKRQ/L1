package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 {
	if p.x == other.x && p.y == other.y {
		return 0
	}
	if p.x == other.x {
		return math.Abs(p.y - other.y)
	}
	if p.y == other.y {
		return math.Abs(p.x - other.x)
	}
	return math.Sqrt(math.Pow(math.Abs(p.x-other.x), 2) + math.Pow(math.Abs(p.y-other.y), 2))
}

func main() {
	p1 := NewPoint(3.0, 5.0)
	p2 := NewPoint(7.0, 2.0)
	fmt.Println(p1.Distance(p2))
}
