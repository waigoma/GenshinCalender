package resin

import (
	"github.com/waigoma/GenshinCalender/src/genshin/character"
	"github.com/waigoma/GenshinCalender/src/genshin/region"
	"github.com/waigoma/GenshinCalender/src/genshin/talent"
)

type Mode int

const (
	// RegenMinute 樹脂が回復する時間
	RegenMinute = 8

	// ModeSecond 算出モード：秒
	ModeSecond Mode = iota
	// ModeMinute 算出モード：分
	ModeMinute
	// ModeHour 算出モード：時
	ModeHour
	// ModeDay 算出モード：日
	ModeDay
)

func CalculateRegenTime(reqResin int, mode Mode) float64 {
	switch mode {
	case ModeSecond:
		return float64(reqResin * RegenMinute * 60)
	case ModeMinute:
		return float64(reqResin * RegenMinute)
	case ModeHour:
		return float64(reqResin * RegenMinute / 60)
	case ModeDay:
		return float64(reqResin * RegenMinute / 60 / 24)
	default:
		return -1
	}
}

func CalculateTotalResin(characterStatList []character.Stats) int {
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
		totalResin += CalculateNeedResin(grindTime)
	}
	return totalResin
}
