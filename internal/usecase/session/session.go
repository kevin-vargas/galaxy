package session

import (
	"github.com/kevin-vargas/galaxy/internal/entity"
	"github.com/kevin-vargas/go-core/cache"
)

type Session struct {
	cache.Cache[entity.SatelliteElems]
}

func New() *Session {
	return &Session{
		Cache: cache.New[entity.SatelliteElems](),
	}
}
