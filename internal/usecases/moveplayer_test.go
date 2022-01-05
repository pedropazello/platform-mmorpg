package usecases_test

import (
	"errors"
	"platformmmorpg/internal/entities"
	"platformmmorpg/internal/usecases"
	"platformmmorpg/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestMovePlayerToInvalidPosition(t *testing.T) {
	repository := &mocks.IPlayerRepository{}
	usecase := usecases.NewMovePlayerUseCase(repository)
	playerId := "playerId"
	_, err := usecase.Execute(playerId, "test")

	expectedError := errors.New("invalid direction")

	if err.Error() != expectedError.Error() {
		t.Errorf("MovePlayer was incorrect, got: %d, want: %d.", err, expectedError)
	}
}

func TestMovePlayerToRightPosition(t *testing.T) {
	repository := &mocks.IPlayerRepository{}
	usecase := usecases.NewMovePlayerUseCase(repository)
	playerId := "playerId"

	oldStatePlayer := entities.Player{
		CurrentPositionX: 1,
		CurrentPositionY: 1,
	}
	repository.On("Get", playerId).Return(&oldStatePlayer, nil)
	repository.On("Save", mock.Anything).Return(nil)

	newStatePlayer, err := usecase.Execute(playerId, "right")

	expectedPositionX := 2

	if err != nil {
		t.Errorf("MovePlayer was incorrect, got: %d", err)
	}

	if newStatePlayer.CurrentPositionX != expectedPositionX {
		t.Errorf("MovePlayer was incorrect, got: %d, want: %d.",
			newStatePlayer.CurrentPositionX, expectedPositionX)
	}
}

func TestMovePlayerToLeftPosition(t *testing.T) {
	repository := &mocks.IPlayerRepository{}
	usecase := usecases.NewMovePlayerUseCase(repository)
	playerId := "playerId"

	oldStatePlayer := entities.Player{
		CurrentPositionX: 2,
		CurrentPositionY: 1,
	}
	repository.On("Get", playerId).Return(&oldStatePlayer, nil)
	repository.On("Save", mock.Anything).Return(nil)

	newStatePlayer, err := usecase.Execute(playerId, "left")

	expectedPositionX := 1

	if err != nil {
		t.Errorf("MovePlayer was incorrect, got: %d", err)
	}

	if newStatePlayer.CurrentPositionX != expectedPositionX {
		t.Errorf("MovePlayer was incorrect, got: %d, want: %d.",
			newStatePlayer.CurrentPositionX, expectedPositionX)
	}
}
