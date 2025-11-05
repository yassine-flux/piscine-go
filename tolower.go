package piscine

func ToLower(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			char += 32
			result += string(char)
		} else {
			result += string(char)
		}
	}
	return result
}
