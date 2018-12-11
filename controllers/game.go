package controllers

import (
	"github.com/astaxie/beego"
	"github.com/minesweeper/services"
	"github.com/minesweeper/models/interfaces"
	"github.com/minesweeper/models/domain"
	"encoding/json"
	"github.com/minesweeper/models/commands"
)

var service interfaces.IGameService

func init() {
	service = services.GameServiceInstance()
}

type GamesController struct {
	beego.Controller
}

// @Title Get game by email
// @Summary getGames
// @Description A GET request is used to retrieve a game using an email
// @Param   email path    string  true    "The player email." username@email.com
// @Success 200 {object} domain.Game
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router / [get]
func (c *GamesController) Get() {
	var email string
	c.Ctx.Input.Bind(&email, "email")

	games, err := service.GetGamesByEmail(email)

	if err != nil {
		c.Data["json"] = &err
		c.ServeJSON()
		return
	}

	c.Data["json"] = &games
	c.ServeJSON()
}

// @Title New game
// @Summary postGames
// @Description A POST request is used to create a new game using an email
// @Param   email body    string  true    "The player email." username@email.com
// @Param   size body    int  true    "The size of the board." 10
// @Success 200 {object} domain.Game
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router / [post]
func (c *GamesController) Post() {
	var body commands.CreateGameCommand

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &body)

	if err != nil {
		c.Data["json"] = &err
		c.ServeJSON()
		return
	}

	game := domain.Game{
		Email: body.Email,
		Board: domain.Board{
			Rows: body.Size,
			Cols: body.Size,
		},
	}

	newGame, err := service.StartGame(game)

	if err != nil {
		c.Data["json"] = &err
		c.ServeJSON()
		return
	}

	c.Data["json"] = &newGame
	c.ServeJSON()
}

// @Title Update game
// @Summary putGames
// @Description A PUT request is used to save the game
// @Param   game body    string  true    "The game." -
// @Success 200 {object} domain.Game
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router / [put]
func (c *GamesController) Put() {
}

// @Title Open Cell
// @Summary patchGame
// @Description A PATCH request is used to open the squares
// @Param   gameId body    string  true    "The game's id." 1
// @Param   email body    string  true    "The player's email." 1
// @Param   cell body    string  true    "The cell that will be opened." 1
// @Success 200 {object} domain.Game
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router /open [patch]
func (c *GamesController) PatchOpen() {
	var body commands.ClickCellCommand

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &body)

	if err != nil {
		c.Data["json"] = &err
		c.ServeJSON()
		return
	}

	game, err := service.OpenCell(body.GameId, body.Email, body.Cell)

	if err != nil {
		c.Data["json"] = &err
		c.ServeJSON()
		return
	}

	c.Data["json"] = &game
	c.ServeJSON()
}

// @Title Flag cell
// @Summary patchGame
// @Description A PATCH request is used to flag the squares
// @Param   gameId body    string  true    "The game's id." 1
// @Param   email body    string  true    "The player's email." 1
// @Param   cell body    string  true    "The cell that will be flagged." 1
// @Success 200 {object} domain.Game
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router /flag [patch]
func (c *GamesController) PatchFlag() {
	var body commands.ClickCellCommand

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &body)

	if err != nil {
		c.Data["json"] = &err
		c.ServeJSON()
		return
	}

	game, err := service.FlagCell(body.GameId, body.Email, body.Cell)

	if err != nil {
		c.Data["json"] = &err
		c.ServeJSON()
		return
	}

	c.Data["json"] = &game
	c.ServeJSON()
}
