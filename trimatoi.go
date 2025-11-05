package piscine

func TrimAtoi(s string) int {
	str := ""
	num := 0
	sign := 1
	for _, char := range s {
		if IsNum(char) || char == '-' {
			str += string(char)
		}
	}
	if len(str) > 0 {
		if str[0] == '-' {
			sign = -1
		}
	}
	for _, char := range str {
		if char != '-' {
			num = num*10 + int(char-'0')
		}
	}
	return num * sign
}

func IsNum(r rune) bool {
	return r >= '0' && r <= '9'
}
