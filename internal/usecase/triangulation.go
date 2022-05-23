package usecase

import (
	"errors"

	"github.com/kevin-vargas/galaxy/internal/entity"
)

const (
	MIN_SIZE = 2
)

type triangulationUseCase struct {
	finder     Finder
	satellites Satellites
}

func NewTriangulation(s Satellites, f Finder) Triangulation {
	return &triangulationUseCase{
		finder:     f,
		satellites: s,
	}
}

func (t *triangulationUseCase) GetPoints(positions ...*entity.Position) ([]*entity.Point, error) {
	if len(positions) < MIN_SIZE {
		return nil, errors.New("cannot calculate position")
	}
	circles := make([]*entity.Circle, len(positions))
	for index, pos := range positions {
		satellite, err := t.satellites.GetSatellite(pos.SatelliteName)
		if err != nil {
			return nil, err
		}
		circles[index] = entity.NewCircle(satellite.Point, pos.Distance)
	}
	intersections, err := makeIntersections(circles)
	if err != nil {
		return nil, err
	}
	points, err := t.finder.GetPoints(intersections)
	if err != nil {
		return nil, err
	}
	if len(points) == 0 {
		return nil, errors.New("invalid intersections")
	}
	return points, nil
}

func (t *triangulationUseCase) GetMessage(messages ...[]string) (string, error) {
	msg := entity.NewMessage()
	for _, message := range messages {
		_, err := msg.Add(message)
		if err != nil {
			return "", err
		}
	}
	return msg.Build()
}

func makeIntersections(circles []*entity.Circle) ([][]*entity.Point, error) {
	size := len(circles)
	if size < MIN_SIZE {
		return nil, errors.New("need two elements for intersection")
	}
	result := make([][]*entity.Point, 0)
	first := circles[0]
	for i := 1; i < size; i++ {
		elem := circles[i]
		result = append(result, first.Intersect(elem))
	}
	return result, nil
}
