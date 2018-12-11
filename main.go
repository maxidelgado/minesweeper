package main

import (
	_ "github.com/minesweeper/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/api/v1/swagger"] = "swagger"
	}

	beego.Run()
}

