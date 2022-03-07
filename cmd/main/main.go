package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/internal/genshin/character"
	"github.com/waigoma/GenshinCalender/internal/genshin/talent"
	"github.com/waigoma/GenshinCalender/internal/yml"
	"github.com/waigoma/GenshinCalender/web"
)

func main() {
	loadYml()

	router := gin.Default()
	router.Static("/static", "web/static")
	router.LoadHTMLGlob("web/template/*.html")

	web.RegisterIndexHandler(router)
	web.RegisterResultHandler(router)

	err := router.Run()

	if err != nil {
		return
	}

}

func loadYml() {
	character.Initialize(yml.LoadCharacters())
	talent.InitializeBook(yml.LoadTalentBooks())
	talent.InitializeBookCount(yml.LoadTalentBookCounts())
}
