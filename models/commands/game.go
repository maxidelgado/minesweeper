package commands

import "github.com/minesweeper/models/domain"

type CreateGameCommand struct {
	Email string `json:"email"`
	Size  int    `json:"size"`
}

type ClickCellCommand struct {
	GameId int         `json:"gameId"`
	Email  string      `json:"email"`
	Cell   domain.Cell `json:"cell"`
}
