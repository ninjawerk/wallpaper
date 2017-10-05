package controllers

import (
	"github.com/BeastSanchez/wallpaper/models"
	"github.com/astaxie/beego/utils/pagination"
)

type CategoriesController struct {
	BaseController
}

func (c *CategoriesController) Get() {
	c.TplName = "categories.tpl"
	postsPerPage :=20
	q:=c.GetString("id")
	total, _ :=models.GetFromTagCount(q)
	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, int64(total) )
	c.Data["posts"] ,_  = models.GetFromTag(q,paginator.Offset(),postsPerPage)
	c.Data["pageTitle"] = "Wallpapers for " + q
	c.Data["postTotal"] = total
}
