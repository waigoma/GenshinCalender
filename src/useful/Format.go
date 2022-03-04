package useful

import "strconv"

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
