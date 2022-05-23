package usecase

import "github.com/kevin-vargas/galaxy/internal/entity"

type (
	Triangulation interface {
		GetPoints(positions ...*entity.Position) ([]*entity.Point, error)
		GetMessage(...[]string) (string, error)
	}

	Satellites interface {
		GetSatellite(s string) (*entity.Satellite, error)
	}

	Finder interface {
		GetPoints([][]*entity.Point) ([]*entity.Point, error)
	}
)

type (
	Session interface {
		Get(string) (entity.SatelliteElems, bool)
		Set(string, entity.SatelliteElems)
	}
)
