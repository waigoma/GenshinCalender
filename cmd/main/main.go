package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/internal/genshin/character"
	talent2 "github.com/waigoma/GenshinCalender/internal/genshin/talent"
	yml2 "github.com/waigoma/GenshinCalender/internal/yml"
	web2 "github.com/waigoma/GenshinCalender/web"
)

func main() {
	loadYml()

	router := gin.Default()
	router.Static("/static", "web/static")
	router.LoadHTMLGlob("web/template/*.html")

	web2.RegisterIndexHandler(router)
	web2.RegisterResultHandler(router)

	err := router.Run()

	if err != nil {
		return
	}

}

func loadYml() {
	character.Initialize(yml2.LoadCharacters())
	talent2.InitializeBook(yml2.LoadTalentBooks())
	talent2.InitializeBookCount(yml2.LoadTalentBookCounts())
}
