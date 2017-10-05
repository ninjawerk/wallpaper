package routers

import (
	"github.com/BeastSanchez/wallpaper/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/category", &controllers.CategoriesController{})
    beego.Router("/wallpaper", &controllers.WallpaperController{})
    beego.Router("/recent", &controllers.RecentController{})
    beego.Router("/berzerkzerglongstringsabcdesanchez", &controllers.UploadController{})
}
