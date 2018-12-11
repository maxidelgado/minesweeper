package controllers

import (
	"github.com/astaxie/beego"
)

type GamesController struct {
	beego.Controller
}

// @Title Get game by email
// @Summary getGames
// @Description A GET request is used to retrieve a game using an email
// @Param   email path    string  true    "The player email." username@email.com
// @Success 200 {object} models.domain.Game
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router /:email [get]
func (c *GamesController) Get() {
}

// @Title New game
// @Summary postGames
// @Description A POST request is used to create a new game using an email
// @Param   email body    string  true    "The player email." username@email.com
// @Param   size body    int  true    "The size of the board." 10
// @Success 200 {object} models.domain.Game
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router / [post]
func (c *GamesController) Post() {
}

// @Title Update game
// @Summary putGames
// @Description A PUT request is used to save the game
// @Param   game body    string  true    "The game." -
// @Success 200 {object} models.domain.Game
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router / [put]
func (c *GamesController) Put() {
}

// @Title Open Cell
// @Summary patchGame
// @Description A PATCH request is used to open the squares
// @Param   id body    string  true    "The game's id." 1
// @Param   square body    string  true    "The square that will be opened." 1
// @Success 200 {object} models.domain.Game
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router /open [patch]
func (c *GamesController) PatchOpen() {
}

// @Title Flag cell
// @Summary patchGame
// @Description A PATCH request is used to flag the squares
// @Param   id body    string  true    "The game's id." 1
// @Param   square body    string  true    "The square that will be flagged." 1
// @Success 200 {object} models.domain.Game
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router /flag [patch]
func (c *GamesController) PatchFlag() {
}
