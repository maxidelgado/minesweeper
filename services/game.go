package services

import (
	"github.com/minesweeper/models/domain"
	"sync"
)

var service GameService
var serviceSingleteon sync.Once

type GameService struct {
	Store map[string]domain.Game
}

func GetGameService() *GameService {
	serviceSingleteon.Do(func() {
		store := make(map[string]domain.Game)
		service = GameService{
			Store: store,
		}
	})

	return &service
}

func (service *GameService) StartGame(game domain.Game) (domain.Game, error) {
	game.Board.Setup()
	service.Store[game.Email] = game

	return service.Store[game.Email], nil
}
func (service *GameService) OpenCell(cell domain.Cell) error {
	return nil
}
