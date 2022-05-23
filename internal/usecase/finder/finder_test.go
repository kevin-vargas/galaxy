package finder

import (
	"testing"

	"github.com/kevin-vargas/galaxy/internal/entity"
	"github.com/stretchr/testify/assert"
)

func Test_GetPoints(t *testing.T) {
	testCases := []struct {
		desc   string
		points [][]*entity.Point
		expect int
	}{
		{
			"simple",
			[][]*entity.Point{{entity.NewPoint(1, 1), entity.NewPoint(2, 2), entity.NewPoint(3, 3)}, {entity.NewPoint(1, 1)}},
			1,
		},
		{
			"dual",
			[][]*entity.Point{{entity.NewPoint(3, 3), entity.NewPoint(1, 1), entity.NewPoint(2, 2)}, {entity.NewPoint(2, 2), entity.NewPoint(1, 1)}},
			2,
		},
		{
			"no intersection",
			[][]*entity.Point{{entity.NewPoint(3, 3), entity.NewPoint(1, 1), entity.NewPoint(2, 2)}, {entity.NewPoint(4, 4)}},
			0,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			// arrange
			finder := New()
			// act
			result, err := finder.GetPoints(tt.points)
			if err != nil {
				t.FailNow()
			}
			// assert
			assert.Len(t, result, tt.expect)
		})
	}
}

func Test_Includes(t *testing.T) {
	testCases := []struct {
		desc   string
		points []*entity.Point
		point  *entity.Point
		expect bool
	}{
		{
			"simple",
			[]*entity.Point{entity.NewPoint(1, 1), entity.NewPoint(2, 2), entity.NewPoint(3, 3)},
			entity.NewPoint(1, 1),
			true,
		},
		{
			"simple in another order",
			[]*entity.Point{entity.NewPoint(3, 3), entity.NewPoint(1, 1), entity.NewPoint(2, 2)},
			entity.NewPoint(2, 2),
			true,
		},
		{
			"simple",
			[]*entity.Point{entity.NewPoint(1, 1), entity.NewPoint(2, 2), entity.NewPoint(3, 3)},
			entity.NewPoint(4, 4),
			false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			// act
			result := includes(tt.points, tt.point)

			// assert
			assert.Equal(t, tt.expect, result)
		})
	}
}
