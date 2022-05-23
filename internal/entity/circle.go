package entity

import "math"

type Circle struct {
	Center *Point
	Radius float64
}

func NewCircle(center *Point, radius float64) *Circle {
	return &Circle{
		Center: center,
		Radius: radius,
	}
}

func (c *Circle) Intersect(circle *Circle) []*Point {
	p0 := c.Center
	p1 := circle.Center

	r0 := c.Radius
	r1 := circle.Radius

	distance := p0.Distance(p1)
	Points := make([]*Point, 2)

	a := (math.Pow(r0, 2) - math.Pow(r1, 2) + math.Pow(distance, 2)) / (2 * distance)
	h := math.Sqrt(math.Pow(r0, 2) - math.Pow(a, 2))
	// calculate P2

	p2 := p1.
		Subtract(p0).
		ScalarMultiply(a / distance).
		Add(p0)

	xValue := h * (p1.Y - p0.Y) / distance
	yValue := h * (p1.X - p0.X) / distance

	Points[0] = &Point{
		X: p2.X + xValue,
		Y: p2.Y - yValue,
	}

	Points[1] = &Point{
		X: p2.X - xValue,
		Y: p2.Y + yValue,
	}

	return Points
}
