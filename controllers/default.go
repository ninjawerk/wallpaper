package controllers

import (
	"github.com/VoidArtanis/wallpaper/models"
	"sort"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	//cats7 :=  new(models.Categories )
	//cats7.Data = []models.Category{
	//	{
	//		Title : "McCree",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/mccree/hero-select-portrait.png",
	//	},{
	//		Title : "Genji",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/genji/hero-select-portrait.png",
	//	},
	//	{
	//		Title : "Pharah",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/pharah/hero-select-portrait.png",
	//	},
	//	{
	//		Title : "Reaper",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/reaper/hero-select-portrait.png",
	//	},{
	//		Title : "Soldier 76",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/soldier-76/hero-select-portrait.png",
	//	},{
	//		Title : "Sombra",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/sombra/hero-select-portrait.png",
	//	},{
	//		Title : "Tracer",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/tracer/hero-select-portrait.png",
	//	},{
	//		Title : "Bastion",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/bastion/hero-select-portrait.png",
	//	},{
	//		Title : "Hanzo",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/hanzo/hero-select-portrait.png",
	//	},{
	//		Title : "Junkrat",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/junkrat/hero-select-portrait.png",
	//	},{
	//		Title : "Mei",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/mei/hero-select-portrait.png",
	//	},{
	//		Title : "Torbjorn",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/torbjorn/hero-select-portrait.png",
	//	},{
	//		Title : "Widowmaker",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/widowmaker/hero-select-portrait.png",
	//	},{
	//		Title : "D. Va",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/dva/hero-select-portrait.png",
	//	},{
	//		Title : "Orisa",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/orisa/hero-select-portrait.png",
	//	},{
	//		Title : "Reinhardt",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/reinhardt/hero-select-portrait.png",
	//	},{
	//		Title : "Roadhog",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/roadhog/hero-select-portrait.png",
	//	},{
	//		Title : "Winston",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/winston/hero-select-portrait.png",
	//	},{
	//		Title : "Zarya",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/zarya/hero-select-portrait.png",
	//	},{
	//		Title : "Ana",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/ana/hero-select-portrait.png",
	//	},{
	//		Title : "Lucio",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/lucio/hero-select-portrait.png",
	//	},{
	//		Title : "Mercy",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/mercy/hero-select-portrait.png",
	//	},{
	//		Title : "Symmetra",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/symmetra/hero-select-portrait.png",
	//	},{
	//		Title : "Zenyatta",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/zenyatta/hero-select-portrait.png",
	//	},{
	//		Title : "Doomfist",
	//		Description: "ex",
	//		Thumb: "https://blzgdapipro-a.akamaihd.net/hero/doomfist/hero-select-portrait.png",
	//	},
	//}
	//cats7.Save()
	cats ,_ := models.GetAllCategories()
	sort.Slice(cats.Data, func(i, j int) bool {
		return cats.Data[i].Title < cats.Data[j].Title
	})
	c.Data["cats"] = cats.Data
	c.TplName = "home.tpl"
}
