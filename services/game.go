package services

import (
	"github.com/minesweeper/models/domain"
	"sync"
	"github.com/pkg/errors"
)

var service GameService
var serviceSingleteon sync.Once

type GameService struct {
	Store map[string][]domain.Game
}

func GameServiceInstance() *GameService {
	serviceSingleteon.Do(func() {
		store := make(map[string][]domain.Game)
		service = GameService{
			Store: store,
		}
	})

	return &service
}

func (service *GameService) StartGame(game domain.Game) (domain.Game, error) {
	game.Id = len(service.Store[game.Email]) + 1
	game.Board.Setup()
	service.Store[game.Email] = append(service.Store[game.Email], game)

	return game, nil
}

func (service *GameService) OpenCell(gameId int, email string, cell domain.Cell) (domain.Game, error) {
	for i, game := range service.Store[email] {
		if game.Id == gameId {
			game.Board.OpenCell(cell)
			service.Store[email][i] = game
			return service.Store[email][i], nil
		}
	}

	return domain.Game{}, errors.New("Game not found.")
}

func (service *GameService) FlagCell(gameId int, email string, cell domain.Cell) (domain.Game, error) {
	for i, game := range service.Store[email] {
		if game.Id == gameId {
			game.Board.FlagCell(cell)
			service.Store[email][i] = game
			return service.Store[email][i], nil
		}
	}

	return domain.Game{}, errors.New("Game not found.")
}

func (service *GameService) GetGamesByEmail(email string) ([]domain.Game, error) {
	if len(service.Store[email]) == 0 {
		return nil, errors.New("No content.")
	}

	return service.Store[email], nil
}
