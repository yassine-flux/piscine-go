package piscine

func ToUpper(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char -= 32
			result += string(char)
		} else {
			result += string(char)
		}
	}
	return result
}
