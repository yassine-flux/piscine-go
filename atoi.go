package piscine

func Atoi(s string) int {
	sign := 1
	num := 0
	start := 0
	if len(s) == 0 {
		return 0
	}
	if s[0] == '-' {
		start = 1
		sign = -1
	}
	if s[0] == '+' {
		start = 1
	}
	for i := start; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		} else {
			return 0
		}
	}
	num *= sign
	return num
}
