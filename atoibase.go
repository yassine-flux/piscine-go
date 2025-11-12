package main

import "fmt"

func AtoiBase(s string, base string) int {
	if iSValidBase(base) {
		num := Atoi(s)

		



		return num
	} else {
		return 0
	}
}

func iSValidBase(s string) bool {
	if len(s) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, char := range s {
		if char == '+' || char == '-' {
			return false
		}
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}


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

	func main() {
		fmt.Println(AtoiBase("125", "0123456789"))
		fmt.Println(AtoiBase("1111101", "01"))
		fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
		fmt.Println(AtoiBase("uoi", "choumi"))
		fmt.Println(AtoiBase("bbbbbab", "-ab"))
	}