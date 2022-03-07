package talent

// CalculateTalentBooks common 換算でどれくらい必要かを計算する
func CalculateTalentBooks(bookCount map[string]int) int {
	var books int
	for key, value := range bookCount {
		switch key {
		case "common":
			books += value
		case "rare":
			books += value * 3
		case "epic":
			books += value * 9
		}
	}
	return books
}
