package web

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/src/genshin/character"
	"github.com/waigoma/GenshinCalender/src/genshin/resin"
	"github.com/waigoma/GenshinCalender/src/genshin/talent"
	"github.com/waigoma/GenshinCalender/src/useful"
	"net/http"
)

type Type int

type SelectForm struct {
	CharacterName string
	TalentForms   []TalentForm
}

type TalentForm struct {
	Type Type
	From string
	To   string
}

type DropForm struct {
	Common string
	Rare   string
	Epic   string
}

const (
	NormalAttack Type = iota
	SkillAttack
	ULTAttack
)

var dropForm DropForm

func RegisterResultHandler(router *gin.Engine) {
	router.GET("/result", resultGetHandle)
	router.POST("/result", resultPostHandle)
}

func resultGetHandle(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/")
	ctx.HTML(
		http.StatusOK,
		"result.html",
		gin.H{
			"selectedCharacters": character.GetAllCharacters(),
		})
}

func resultPostHandle(ctx *gin.Context) {
	// フォームから受け取ったデータを CharacterStat に変換
	characterStatList := getResult(initSelectForm(ctx))

	// 必要樹脂数
	customDrop := map[string]int{
		"common": useful.StringToInt(dropForm.Common),
		"rare":   useful.StringToInt(dropForm.Rare),
		"epic":   useful.StringToInt(dropForm.Epic),
	}

	totalResin := resin.CalculateTotalResin(characterStatList, customDrop)

	// 回復時間 (分)
	totalTime := resin.CalculateRegenTime(totalResin, resin.ModeMinute)

	totalTimeStr := useful.MinuteToTime(int(totalTime))

	ctx.HTML(
		http.StatusOK,
		"result.html",
		gin.H{
			"characterStatList": characterStatList,
			"totalResin":        totalResin,
			"condensedResin":    totalResin / 40,
			"totalTime":         totalTimeStr,
		})
}

// フォームから受け取ったデータを処理する
func initSelectForm(ctx *gin.Context) []SelectForm {
	var fromArray [][]string
	var toArray [][]string

	// post された情報取得
	for idx, tmp := range []string{"from", "to"} {
		for _, num := range []string{"1", "2", "3"} {
			if idx == 0 {
				fromArray = append(fromArray, ctx.PostFormArray(tmp+num))
			} else if idx == 1 {
				toArray = append(toArray, ctx.PostFormArray(tmp+num))
			}
		}
	}

	// 必要なデータを Context から取得
	characterArray := ctx.PostFormArray("character")
	NormalAttackArray := ctx.PostFormArray("normalAttack")
	SkillAttackArray := ctx.PostFormArray("skillAttack")
	ULTAttackArray := ctx.PostFormArray("ultAttack")

	// SelectForm struct にプロットしていく
	var selectForms []SelectForm

	for idx, characterName := range characterArray {
		var selectForm SelectForm

		selectForm.CharacterName = characterName

		if useful.Contains(NormalAttackArray, characterName) {
			selectForm.TalentForms = append(selectForm.TalentForms, TalentForm{
				Type: NormalAttack,
				From: fromArray[0][idx],
				To:   toArray[0][idx],
			})
		}

		if useful.Contains(SkillAttackArray, characterName) {
			selectForm.TalentForms = append(selectForm.TalentForms, TalentForm{
				Type: SkillAttack,
				From: fromArray[1][idx],
				To:   toArray[1][idx],
			})
		}

		if useful.Contains(ULTAttackArray, characterName) {
			selectForm.TalentForms = append(selectForm.TalentForms, TalentForm{
				Type: ULTAttack,
				From: fromArray[2][idx],
				To:   toArray[2][idx],
			})
		}

		selectForms = append(selectForms, selectForm)
	}

	// 必要なデータを Context から取得
	commonBook := ctx.PostForm("common")
	rareBook := ctx.PostForm("rare")
	epicBook := ctx.PostForm("epic")

	dropForm.Common = commonBook
	dropForm.Rare = rareBook
	dropForm.Epic = epicBook

	return selectForms
}

// html に渡す値を作成
func getResult(selectForms []SelectForm) []character.Stats {
	var characterStatList []character.Stats

	// 選択したキャラクターをすべてカウント
	for _, selectForm := range selectForms {
		// 選択したキャラクター
		chara := character.GetCharacter(selectForm.CharacterName)

		// 天賦本の数
		talentBookCount := make(map[string]int)

		for _, talentForm := range selectForm.TalentForms {
			bookCounts := talent.CountTalentBooks(talentForm.From, talentForm.To)

			for key, value := range bookCounts {
				if v, ok := talentBookCount[key]; ok {
					talentBookCount[key] = v + value
				} else {
					talentBookCount[key] = value
				}
			}
		}

		// 天賦本を取得
		talentBook := talent.GetTalentBook(chara.TalentBook)

		// 天賦本名と数保存用
		var talents []character.Talent

		// 天賦本の数を取得
		for key, value := range talentBookCount {
			// 天賦レアリティ名と一致した場合
			if val, ok := talentBook.RarityName[key]; ok {
				talents = append(talents, character.Talent{Type: key, Name: val, Count: value})
			}
		}

		characterStatList = append(characterStatList, character.Stats{Character: chara, Talent: talents, Day: talentBook.Day})
	}

	return characterStatList
}
