package web

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/internal/genshin/character"
	"net/http"
)

func RegisterIndexHandler(router *gin.Engine) {
	router.GET("/", indexGetHandle)
	router.POST("/", indexPostHandle)
}

func indexGetHandle(ctx *gin.Context) {
	characters := character.GetAllCharacters()

	// region map
	var characterMap = make(map[string][]character.Character)

	for _, chara := range characters {
		if cs, ok := characterMap[chara.JPNation]; ok {
			characterMap[chara.JPNation] = append(cs, chara)
		} else {
			characterMap[chara.JPNation] = []character.Character{chara}
		}
	}

	// sort by region
	var keys = []string{"モンド", "璃月", "稲妻", "スメール"}
	var characterMapList []map[string][]character.Character

	for _, key := range keys {
		characterMapList = append(characterMapList, map[string][]character.Character{key: characterMap[key]})
	}

	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"characters": characterMapList,
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
