package web

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/src/yml"
	"net/http"
)

var characters yml.Characters
var talentBooks yml.TalentBooks
var talentBookCounts yml.TalentBookCounts

func RegisterIndexHandler(router *gin.Engine, chs yml.Characters, tbs yml.TalentBooks, tbc yml.TalentBookCounts) {
	characters = chs
	talentBooks = tbs
	talentBookCounts = tbc
	router.GET("/", indexGetHandle)
	router.POST("/", indexPostHandle)
}

func indexGetHandle(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"characters": characters.Characters,
		})
}

func indexPostHandle(ctx *gin.Context) {
	selectedCharacters := ctx.PostFormArray("selectCharacter")
	var selectedCharactersList []yml.CharacterStats

	// 選択したキャラクター
	for _, sc := range selectedCharacters {
		// データとして用意されたキャラクター
		for _, chara := range characters.Characters {

			// 名前一致しない場合はスキップ
			if chara.ENName != sc {
				continue
			}

			// 先にレベル上げの数を取得 common : 3
			var talentBookCount map[string]int
			for _, tbc := range talentBookCounts.TalentBookCounts {
				if tbc.NtoM == "1-10" {
					talentBookCount = tbc.BookCount
				}
			}

			// 天賦本を取得
			for _, tb := range talentBooks.TalentBooks {
				// 一致で登録
				if tb.ENName == chara.TalentBook {
					// 天賦本名と数を登録
					var talent []yml.Talent
					// 天賦本の数を取得
					for key, value := range talentBookCount {
						// 天賦本名と一致した場合
						if val, ok := tb.RarityName[key]; ok {
							talent = append(talent, yml.Talent{Name: val, Count: value})
						}
					}
					selectedCharactersList = append(selectedCharactersList, yml.CharacterStats{Character: chara, Talent: talent})
				}
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
