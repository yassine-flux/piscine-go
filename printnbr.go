package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		n = 223372036854775808
	}

	if n < 0 {
		n = -n
		z01.PrintRune('-')
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	digits := []rune{}
	for n > 0 {
		digit := n % 10
		digits = append(digits, rune(digit)+'0')
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
