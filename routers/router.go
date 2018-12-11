package routers

import (
	"github.com/minesweeper/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

// @APIVersion 1.0.0
// @Title Minesweeper API
// @Description Web Minesweeper challenge for Fukuroo!
// @Contact maxi.delga2@gmail.com
func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET"},
	}))

	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/games",
			beego.NSInclude(
				&controllers.GamesController{},
			)),
	)

	beego.AddNamespace(ns)
}
