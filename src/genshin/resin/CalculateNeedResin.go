package resin

const (
	// PerOnce 秘境を 1 回行くごとに消費する樹脂数
	PerOnce = 20
)

func CalculateNeedResin(grindTime int) int {
	return grindTime * PerOnce
}
