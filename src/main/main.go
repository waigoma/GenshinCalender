package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/src/web"
	"github.com/waigoma/GenshinCalender/src/yml"
)

var characters yml.Characters
var talentBooks yml.TalentBooks
var talentBookCounts yml.TalentBookCounts

func main() {
	characters = yml.LoadCharacters()
	talentBooks = yml.LoadTalentBooks()
	talentBookCounts = yml.LoadTalentBookCounts()

	router := gin.Default()
	router.Static("/src", "templates/src")
	router.LoadHTMLGlob("templates/*.html")

	web.RegisterIndexHandler(router, characters)
	web.RegisterResultHandler(router, characters)

	err := router.Run()

	if err != nil {
		return
	}

}
