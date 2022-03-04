package web

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/src/genshin/character"
	"github.com/waigoma/GenshinCalender/src/genshin/talent"
	"net/http"
)

func RegisterResultHandler(router *gin.Engine) {
	router.GET("/result", resultGetHandle)
	router.POST("/result", resultPostHandle)
}

func resultGetHandle(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"result.html",
		gin.H{
			"selectedCharacters": character.GetAllCharacters(),
		})
}

func resultPostHandle(ctx *gin.Context) {
	characterArray := ctx.PostFormArray("character")
	fromArray := ctx.PostFormArray("from")
	toArray := ctx.PostFormArray("to")

	characterStatList := getResult(characterArray, fromArray, toArray)

	ctx.HTML(
		http.StatusOK,
		"result.html",
		gin.H{
			"characterStatList": characterStatList,
		})
}

func getResult(characterArray []string, fromArray []string, toArray []string) []character.Stats {
	var characterStatList []character.Stats

	// 選択したキャラクターをすべてカウント
	for idx, characterName := range characterArray {
		// 選択したキャラクター
		chara := character.GetCharacter(characterName)

		// 天賦本の数
		talentBookCount := talent.CountTalentBooks(fromArray[idx], toArray[idx])

		// 天賦本を取得
		talentBook := talent.GetTalentBook(chara.TalentBook)

		// 天賦本名と数保存用
		var talents []character.Talent

		// 天賦本の数を取得
		for key, value := range talentBookCount {
			// 天賦本名と一致した場合
			if val, ok := talentBook.RarityName[key]; ok {
				talents = append(talents, character.Talent{Name: val, Count: value})
			}
		}

		characterStatList = append(characterStatList, character.Stats{Character: chara, Talent: talents})
	}

	return characterStatList
}
