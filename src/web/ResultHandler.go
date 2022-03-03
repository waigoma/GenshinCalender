package web

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/src/yml"
	"net/http"
)

func RegisterResultHandler(router *gin.Engine, chs yml.Characters) {
	characters = chs
	router.GET("/result", resultGetHandle)
	router.POST("/result", resultPostHandle)
}

func resultGetHandle(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"result.html",
		gin.H{
			"selectedCharacters": characters.Characters,
		})
}

func resultPostHandle(ctx *gin.Context) {
	selectedCharacters := ctx.PostFormArray("selectCharacter")
	var selectedCharactersList []yml.Character

	// 選択したキャラクター
	for _, sc := range selectedCharacters {
		// データとして用意されたキャラクター
		for _, chara := range characters.Characters {
			// 名前一致で選択キャラリストに追加
			if chara.ENName == sc {
				selectedCharactersList = append(selectedCharactersList, chara)
			}
		}
	}

	ctx.HTML(
		http.StatusOK,
		"result.html",
		gin.H{
			"selectedCharacters": selectedCharactersList,
		})
}
