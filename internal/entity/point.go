package entity

import (
	"math"
)

const DeltaPoint = 0.001

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (p *Point) Distance(point *Point) float64 {
	return math.Sqrt(math.Pow(p.X-point.X, 2) + math.Pow(p.Y-point.Y, 2))
}

func (p *Point) Subtract(point *Point) *Point {
	return &Point{
		X: p.X - point.X,
		Y: p.Y - point.Y,
	}
}

func (p *Point) Add(point *Point) *Point {
	return &Point{
		X: p.X + point.X,
		Y: p.Y + point.Y,
	}
}

func (p *Point) ScalarMultiply(number float64) *Point {
	return &Point{
		p.X * number,
		p.Y * number,
	}
}

func (p *Point) Equal(point *Point) bool {
	return math.Abs(p.X-point.X) < DeltaPoint && math.Abs(p.Y-point.Y) < DeltaPoint
}
