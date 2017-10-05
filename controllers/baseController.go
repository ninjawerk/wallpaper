package controllers
import (
	"github.com/astaxie/beego"
 )

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	if c.Ctx.Request.Header.Get("X-Pjax") == "" {
		c.Layout="index.tpl"
	}
}

