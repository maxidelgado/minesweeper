package interfaces

import "github.com/minesweeper/models/domain"

type IGameService interface {
	StartGame(game domain.Game) (domain.Game, error)
	OpenCell(gameId int, email string, cell domain.Cell) (domain.Game, error)
	GetGamesByEmail(email string) ([]domain.Game, error)
	FlagCell(gameId int, email string, cell domain.Cell) (domain.Game, error)
}
