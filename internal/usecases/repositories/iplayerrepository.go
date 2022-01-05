package repositories

import "platformmmorpg/internal/entities"

type IPlayerRepository interface {
	Get(playerId string) (*entities.Player, error)
	Save(player *entities.Player) error
}
