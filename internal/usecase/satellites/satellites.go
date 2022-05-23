package satellites

import (
	"errors"

	"github.com/kevin-vargas/galaxy/internal/entity"
)

type Satellites struct{}

func New() *Satellites {
	return &Satellites{}
}

func (s *Satellites) GetSatellite(str string) (*entity.Satellite, error) {
	for _, satellite := range satellites {
		if satellite.Name == str {
			return satellite, nil
		}
	}
	return nil, errors.New("unknown satellite")
}
