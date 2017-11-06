package controllers

import (
	"github.com/VoidArtanis/wallpaper/models"
	"strings"
)

type WallpaperController struct {
	BaseController
}

func (c *WallpaperController) Get() {
	c.TplName = "wallpaper.tpl"
	q:=c.GetString("id")
	obj, _ :=models.GetFromId(q)
	c.Data["post"]  =obj
	c.Data["tags"]  =strings.Join(obj.Tags, ", ")
}
