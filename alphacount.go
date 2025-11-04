package piscine

func AlphaCount(s string) int {
	count := 0
	for _, char := range s {
		if isAlpha(char) {
			count++
		}
	}
	return count
}

func isAlpha(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}
