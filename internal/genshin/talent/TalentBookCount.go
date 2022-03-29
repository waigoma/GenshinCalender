package talent

import (
	"github.com/waigoma/GenshinCalender/pkg/useful"
	"strconv"
)

var talentBookCounts []BookCount

type BookCount struct {
	NtoM      string
	Mora      int
	BookCount map[string]int
}

func InitializeBookCount(talentBookCountsList []BookCount) {
	talentBookCounts = talentBookCountsList
}

func GetTalentBookCount(ntoM string) BookCount {
	for _, talentBookCount := range talentBookCounts {
		if talentBookCount.NtoM == ntoM {
			return talentBookCount
		}
	}

	return BookCount{}
}

func CountTalentBooks(fromStr string, toStr string) map[string]int {
	if fromStr == "1" && toStr == "10" {
		return GetTalentBookCount("1-10").BookCount
	}

	from := useful.StringToInt(fromStr)
	to := useful.StringToInt(toStr)
	fromTo := to - from

	talentBookCount := make(map[string]int)

	for i := 0; i < fromTo; i++ {
		fromLevel := from + i
		toLevel := fromLevel + 1

		// 天賦本の必要数を取得
		count := GetTalentBookCount(strconv.Itoa(fromLevel) + "-" + strconv.Itoa(toLevel))

		// 追加
		for key, value := range count.BookCount {
			if v, ok := talentBookCount[key]; ok {
				talentBookCount[key] = v + value
			} else {
				talentBookCount[key] = value
			}
		}
	}

	return talentBookCount
}
