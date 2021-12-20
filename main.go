package main

import (
	"Beego_Restful_Api/database"
	_ "Beego_Restful_Api/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	database.InitDB()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
