package web

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/src/genshin/character"
	"github.com/waigoma/GenshinCalender/src/genshin/region"
	"github.com/waigoma/GenshinCalender/src/genshin/resin"
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

	var totalResin int

	for _, characterStat := range characterStatList {
		// 必要な天賦本数を計算
		var needTalent int
		for _, books := range characterStat.Talent {
			needTalent += talent.CalculateTalentBooks(map[string]int{books.Type: books.Count})
		}

		// 周回数を計算
		grindTime := region.CalculateGrind(needTalent)

		// 必要な樹脂数
		totalResin += resin.CalculateNeedResin(grindTime)
	}

	// 回復時間 (分)
	totalTime := resin.CalculateRegenTime(totalResin, resin.ModeMinute)

	ctx.HTML(
		http.StatusOK,
		"result.html",
		gin.H{
			"characterStatList": characterStatList,
			"totalResin":        totalResin,
			"totalTime":         totalTime,
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
				talents = append(talents, character.Talent{Type: key, Name: val, Count: value})
			}
		}

		characterStatList = append(characterStatList, character.Stats{Character: chara, Talent: talents})
	}

	return characterStatList
}
