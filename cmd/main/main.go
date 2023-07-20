package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/internal/genshin/character"
	"github.com/waigoma/GenshinCalender/internal/genshin/talent"
	"github.com/waigoma/GenshinCalender/internal/yml"
	"github.com/waigoma/GenshinCalender/web"
	"os"
	"strconv"
)

const isLocal bool = false

func main() {
	port, _ := strconv.Atoi(os.Args[1])
	loadYml()

	router := gin.Default()

	if isLocal {
		// in local
		router.Static("/static", "web/static")
		router.LoadHTMLGlob("web/template/*.html")
	} else {
		// in docker
		router.Static("/static", "/app/web/static")
		router.LoadHTMLGlob("/app/web/template/*.html")
	}
	web.RegisterIndexHandler(router)
	web.RegisterResultHandler(router)

	err := router.Run(":" + strconv.Itoa(port))

	if err != nil {
		return
	}

}

func loadYml() {
	character.Initialize(yml.LoadCharacters(isLocal))
	talent.InitializeBook(yml.LoadTalentBooks(isLocal))
	talent.InitializeBookCount(yml.LoadTalentBookCounts(isLocal))
}
