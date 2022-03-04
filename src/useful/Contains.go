package useful

func Contains(baseStr []string, matchStr string) bool {
	for _, base := range baseStr {
		if base == matchStr {
			return true
		}
	}
	return false
}
