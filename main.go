package main

import (
	_ "github.com/VoidArtanis/wallpaper/routers"
	"github.com/astaxie/beego"
	"github.com/VoidArtanis/wallpaper/shared"
)

func main() {
	projId := beego.AppConfig.String("gds_project_id")
	shared.ConnectGDS(projId)
	shared.ConnectGCS(projId)
	beego.Run()
}

