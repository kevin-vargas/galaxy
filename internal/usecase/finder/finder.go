package finder

import (
	"errors"

	"github.com/kevin-vargas/fp"
	"github.com/kevin-vargas/galaxy/internal/entity"
)

type Finder struct {
}

func New() *Finder {
	return &Finder{}
}
func includes(arr []*entity.Point, p *entity.Point) bool {
	predicate := func(p0 *entity.Point) bool {
		return p.Equal(p0)
	}
	filtered := fp.Filter(arr, predicate)
	return len(filtered) > 0
}

func (f *Finder) GetPoints(points [][]*entity.Point) ([]*entity.Point, error) {
	size := len(points)
	if size == 0 {
		return nil, errors.New("no points")
	}
	first := points[0]
	if size == 1 {
		return first, nil
	}
	reducer := func(acum []*entity.Point, actual []*entity.Point, i int) []*entity.Point {
		newAcum := make([]*entity.Point, 0)
		for _, elem := range acum {
			if includes(actual, elem) {
				newAcum = append(newAcum, elem)
			}
		}
		return newAcum
	}
	return fp.Reduce(points, reducer, first), nil
}
