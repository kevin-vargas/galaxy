// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	entity "github.com/kevin-vargas/galaxy/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// Session is an autogenerated mock type for the Session type
type Session struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *Session) Get(_a0 string) (entity.SatelliteElems, bool) {
	ret := _m.Called(_a0)

	var r0 entity.SatelliteElems
	if rf, ok := ret.Get(0).(func(string) entity.SatelliteElems); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.SatelliteElems)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Set provides a mock function with given fields: _a0, _a1
func (_m *Session) Set(_a0 string, _a1 entity.SatelliteElems) {
	_m.Called(_a0, _a1)
}
