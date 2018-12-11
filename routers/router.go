package routers

import (
	"github.com/minesweeper/controllers"
	"github.com/astaxie/beego"
)

// @APIVersion 1.0.0
// @Title Minesweeper API
// @Description Web Minesweeper challenge for Fukuroo!
// @Contact maxi.delga2@gmail.com
func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/games",
			beego.NSInclude(
				&controllers.GamesController{},
			)),
	)

	beego.AddNamespace(ns)
}