package controllers

import (
	"github.com/VoidArtanis/wallpaper/models"
	"github.com/astaxie/beego/utils/pagination"
)

type RecentController struct {
	BaseController
}

func (c *RecentController) Get() {
	c.TplName = "categories.tpl"
	postsPerPage :=20
	total, _ :=models.GetTotalRecentWallpapers()
	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, int64(total) )
	c.Data["posts"] ,_  = models.GetRecentWallpapers(paginator.Offset(),postsPerPage)
	c.Data["pageTitle"] = "Latest Wallpapers"
	c.Data["postTotal"] = total

}
