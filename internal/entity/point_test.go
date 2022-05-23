package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	Delta = 0.01
)

func Test_Distance(t *testing.T) {

	originPoint := &Point{
		X: 0,
		Y: 0,
	}

	testCases := []struct {
		desc   string
		point  *Point
		expect float64
	}{
		{
			"Different Point",
			&Point{
				X: 1,
				Y: 1,
			},
			1.414214,
		},
		{
			"Same Point",
			&Point{
				X: 0,
				Y: 0,
			},
			0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {

			// act
			result := originPoint.Distance(tt.point)

			// assert
			assert.InDelta(t, tt.expect, result, Delta)
		})
	}
}

func Test_ScalarMultiply(t *testing.T) {
	testCases := []struct {
		desc   string
		point  *Point
		number float64
		expect *Point
	}{
		{
			"Identity element",
			&Point{
				X: 1,
				Y: 1,
			},
			1,
			&Point{
				X: 1,
				Y: 1,
			},
		},
		{
			"Absorving element",
			&Point{
				X: 1,
				Y: 1,
			},
			0,
			&Point{
				X: 0,
				Y: 0,
			},
		},
		{
			"Positive element",
			&Point{
				X: 1,
				Y: 1,
			},
			2,
			&Point{
				X: 2,
				Y: 2,
			},
		},
		{
			"Negative element",
			&Point{
				X: 1,
				Y: 1,
			},
			-2,
			&Point{
				X: -2,
				Y: -2,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {

			// act
			result := tt.point.ScalarMultiply(tt.number)

			// assert
			assert.InDelta(t, tt.expect.X, result.X, Delta)
			assert.InDelta(t, tt.expect.Y, result.Y, Delta)
		})
	}
}

func Test_Add(t *testing.T) {
	testCases := []struct {
		desc   string
		point  *Point
		toAdd  *Point
		expect *Point
	}{
		{
			"Identity element",
			&Point{
				X: 1,
				Y: 1,
			},
			&Point{
				X: 0,
				Y: 0,
			},
			&Point{
				X: 1,
				Y: 1,
			},
		},
		{
			"Positive element",
			&Point{
				X: 1,
				Y: 1,
			},
			&Point{
				X: 1,
				Y: 1,
			},
			&Point{
				X: 2,
				Y: 2,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {

			// act
			result := tt.point.Add(tt.toAdd)

			// assert
			assert.InDelta(t, tt.expect.X, result.X, Delta)
			assert.InDelta(t, tt.expect.Y, result.Y, Delta)
		})
	}
}

func Test_Subtract(t *testing.T) {
	testCases := []struct {
		desc   string
		point  *Point
		toAdd  *Point
		expect *Point
	}{
		{
			"Identity element",
			&Point{
				X: 1,
				Y: 1,
			},
			&Point{
				X: 0,
				Y: 0,
			},
			&Point{
				X: 1,
				Y: 1,
			},
		},
		{
			"Positive element",
			&Point{
				X: 1,
				Y: 1,
			},
			&Point{
				X: 1,
				Y: 1,
			},
			&Point{
				X: 0,
				Y: 0,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {

			// act
			result := tt.point.Subtract(tt.toAdd)

			// assert
			assert.InDelta(t, tt.expect.X, result.X, Delta)
			assert.InDelta(t, tt.expect.Y, result.Y, Delta)
		})
	}
}

func Test_Equal(t *testing.T) {
	testCases := []struct {
		desc   string
		p      *Point
		p0     *Point
		expect bool
	}{
		{
			"Same element",
			&Point{
				X: 1,
				Y: 1,
			},
			&Point{
				X: 1,
				Y: 1,
			},
			true,
		},
		{
			"Different element",
			&Point{
				X: 1,
				Y: 1,
			},
			&Point{
				X: 2,
				Y: 2,
			},
			false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {

			// act
			result := tt.p.Equal(tt.p0)

			// assert
			assert.Equal(t, tt.expect, result)
		})
	}
}
