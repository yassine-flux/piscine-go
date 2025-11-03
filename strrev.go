package piscine

func StrRev(s string) string {
	var result string
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}
