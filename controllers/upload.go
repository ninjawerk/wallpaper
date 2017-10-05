package controllers

import (
	//"golang.org/x/net/context"
	//"google.golang.org/appengine/datastore"
	"github.com/BeastSanchez/wallpaper/models"
	"strings"
	"github.com/BeastSanchez/wallpaper/utils"
)

type UploadController struct {
	BaseController
}

func (c *UploadController) Get() {
	c.TplName = "upload.tpl"
}

func (c *UploadController) Post() {
	c.TplName = "upload.tpl"
	title := c.GetString("title")
	desc := c.GetString("description")
	tags := c.GetString("tags")

	file, _, _ := c.GetFile("file")

	tags = strings.Replace(strings.ToLower(tags), " ", "", -1)
	tagsarr := strings.Split(tags, ",")

	wallpaper := models.Wallpaper{
		Title:       title,
		Description: desc,
		Tags:        tagsarr,
	}

	_ = utils.CreateSizes(file, &wallpaper)
	wallpaper.Save()
}
