package satellites

import "github.com/kevin-vargas/galaxy/internal/entity"

type Type int8

const (
	KENOBIX    = -500
	KENOBIY    = -200
	SKYWALKERX = 100
	SKYWALKERY = -100
	SATOX      = 500
	SATOY      = 100
)

const (
	KENOBI Type = iota - 1
	SKYWALKER
	SATO
)

var satellites map[Type]*entity.Satellite = map[Type]*entity.Satellite{
	KENOBI:    entity.NewSatellite("kenobi", entity.NewPoint(KENOBIX, KENOBIY)),
	SKYWALKER: entity.NewSatellite("skywalker", entity.NewPoint(SKYWALKERX, SKYWALKERY)),
	SATO:      entity.NewSatellite("sato", entity.NewPoint(SATOX, SATOY)),
}

func (t Type) String() string {
	return satellites[t].Name
}
