package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/src/web"
	"github.com/waigoma/GenshinCalender/src/yml"
)

var characters yml.Characters
var talentBooks yml.TalentBooks

func main() {
	characters = yml.LoadCharacters()
	talentBooks = yml.LoadTalentBooks()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	web.RegisterIndexHandler(router, characters)

	err := router.Run()

	if err != nil {
		return
	}

}
