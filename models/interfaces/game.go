package interfaces

import "github.com/minesweeper/models/domain"

type IGameService interface {
	StartGame(game domain.Game) error
	OpenCell(cell domain.Cell) error
}
