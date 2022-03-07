package useful

func ListStringContains(baseStr []string, matchStr string) bool {
	for _, base := range baseStr {
		if base == matchStr {
			return true
		}
	}
	return false
}

func MapStringContains(baseMap map[string]interface{}, matchStr string) bool {
	if _, ok := baseMap[matchStr]; ok {
		return true
	}
	return false
}
