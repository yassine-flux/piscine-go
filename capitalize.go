package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	result := ""
	for i, char := range runes {
		if i == 0 && isAlphabet(char) && isLower(char) {
			char -= 32
		}
		if isAlphabet(char) {
			if isLower(char) && i != 0 && !isAlphabet(runes[i-1]) {
				char -= 32
				result += string(char)
			} else if isUpper(char) && i != 0 && isAlphabet(runes[i-1]) {
				char += 32
				result += string(char)
			} else {
				result += string(char)
			}
		} else {
			result += string(char)
		}
	}
	return result
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isAlphabet(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}
