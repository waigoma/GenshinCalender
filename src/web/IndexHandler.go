package web

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/src/genshin/character"
	"net/http"
)

func RegisterIndexHandler(router *gin.Engine) {
	router.GET("/", indexGetHandle)
	router.POST("/", indexPostHandle)
}

func indexGetHandle(ctx *gin.Context) {
	characters := character.GetAllCharacters()
	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"characters": characters,
		})
}

func indexPostHandle(ctx *gin.Context) {
	selectedCharacters := ctx.PostFormArray("selectedCharacter")
	var charactersList []character.Character

	// 選択したキャラクター
	for _, sc := range selectedCharacters {
		charactersList = append(charactersList, character.GetCharacter(sc))
	}

	ctx.HTML(
		http.StatusOK,
		"select.html",
		gin.H{
			"characters": charactersList,
		})
}
