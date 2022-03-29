package useful

import (
	"fmt"
	"strconv"
	"strings"
)

func MinuteToTime(minutes int) string {
	var time string

	day := minutes / 1440
	hour := (minutes % 1440) / 60
	minute := (minutes % 1440) % 60

	if day > 0 {
		time += strconv.Itoa(day) + " 日 "
	}

	if hour > 0 {
		time += strconv.Itoa(hour) + " 時間 "
	}

	if minute > 0 {
		time += strconv.Itoa(minute) + " 分"
	}

	return time
}

func SplitNumber(price int) string {

	groupingSize := 3
	groupingSeparator := ","

	priceStr := fmt.Sprint(price)
	size := len(priceStr)
	sliceSize := (len(priceStr) + groupingSize - 1) / groupingSize
	priceSlice := make([]string, sliceSize)

	for i := range priceSlice {
		start := size - (sliceSize-i)*groupingSize
		end := start + groupingSize
		if start < 0 {
			start = 0
		}
		priceSlice[i] = priceStr[start:end]
	}

	return strings.Join(priceSlice, groupingSeparator)
}
