package usecases

import (
	"errors"
	"platformmmorpg/internal/entities"
	"platformmmorpg/internal/usecases/repositories"
)

type MovePlayerUseCase struct {
	repository repositories.IPlayerRepository
}

func NewMovePlayerUseCase(repository repositories.IPlayerRepository) *MovePlayerUseCase {
	return &MovePlayerUseCase{
		repository: repository,
	}
}

func (m MovePlayerUseCase) Execute(playerId string, futurePosition string) (*entities.Player, error) {
	validPositions := []string{"left", "right"}

	validPosition := false
	for _, position := range validPositions {
		if position == futurePosition {
			validPosition = true
		}
	}

	if !validPosition {
		return nil, errors.New("invalid direction")
	}

	player, err := m.repository.Get(playerId)
	if err != nil {
		return player, err
	}

	if futurePosition == "right" {
		player.CurrentPositionX++
		m.repository.Save(player)
	}

	if futurePosition == "left" {
		player.CurrentPositionX--
		m.repository.Save(player)
	}

	return player, nil
}
