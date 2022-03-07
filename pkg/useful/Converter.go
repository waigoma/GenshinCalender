package useful

import "strconv"

func StringToInt(str string) int {
	var i, err = strconv.Atoi(str)

	if err != nil {
		return 0
	}

	return i
}
