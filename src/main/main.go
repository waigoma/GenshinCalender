package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/src/genshin/character"
	"github.com/waigoma/GenshinCalender/src/genshin/talent"
	"github.com/waigoma/GenshinCalender/src/web"
	"github.com/waigoma/GenshinCalender/src/yml"
)

func main() {
	loadYml()

	router := gin.Default()
	router.Static("/src", "templates/src")
	router.LoadHTMLGlob("templates/*.html")

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
