package region

const (
	// MinimumGuarantee 最低保証で手に入る天賦本の数
	MinimumGuarantee = 8
)

func CalculateGrind(needBooks int) int {
	return needBooks / MinimumGuarantee
}
