package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/minesweeper/controllers:GamesController"] = append(beego.GlobalControllerRouter["github.com/minesweeper/controllers:GamesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/minesweeper/controllers:GamesController"] = append(beego.GlobalControllerRouter["github.com/minesweeper/controllers:GamesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/minesweeper/controllers:GamesController"] = append(beego.GlobalControllerRouter["github.com/minesweeper/controllers:GamesController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:email`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/minesweeper/controllers:GamesController"] = append(beego.GlobalControllerRouter["github.com/minesweeper/controllers:GamesController"],
		beego.ControllerComments{
			Method: "PatchFlag",
			Router: `/flag`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/minesweeper/controllers:GamesController"] = append(beego.GlobalControllerRouter["github.com/minesweeper/controllers:GamesController"],
		beego.ControllerComments{
			Method: "PatchOpen",
			Router: `/open`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

}
