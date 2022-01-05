// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	entities "platformmmorpg/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// IPlayerRepository is an autogenerated mock type for the IPlayerRepository type
type IPlayerRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: playerId
func (_m *IPlayerRepository) Get(playerId string) (*entities.Player, error) {
	ret := _m.Called(playerId)

	var r0 *entities.Player
	if rf, ok := ret.Get(0).(func(string) *entities.Player); ok {
		r0 = rf(playerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Player)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(playerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: player
func (_m *IPlayerRepository) Save(player *entities.Player) error {
	ret := _m.Called(player)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Player) error); ok {
		r0 = rf(player)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
