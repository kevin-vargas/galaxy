package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Intersect(t *testing.T) {
	testCases := []struct {
		desc        string
		circle      *Circle
		toIntersect *Circle
		expectP0    *Point
		expectP1    *Point
	}{
		{
			"Intersect normal",
			&Circle{
				Center: NewPoint(-1, 0),
				Radius: 3,
			},
			&Circle{
				Center: NewPoint(3, 1),
				Radius: 2,
			},
			NewPoint(1.921, -0.684),
			NewPoint(1.255, 1.978),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {

			// act
			result := tt.circle.Intersect(tt.toIntersect)

			// assert
			p0 := result[0]
			p1 := result[1]
			assert.True(t, p0.Equal(tt.expectP0))
			assert.True(t, p1.Equal(tt.expectP1))
		})
	}
}
