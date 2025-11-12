package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if isValidBase(base) {
		if nbr < 0 {
			z01.PrintRune('-')
			// Handle math.MinInt special case to avoid overflow
			if nbr == -9223372036854775808 {
				// For extreme negative, we'll handle specially
				PrintNbrBase(-(nbr / len(base)), base)
				z01.PrintRune(rune(base[-(nbr % len(base))]))
				return
			}
			nbr = -nbr
		}
		isNegative := false
		if nbr < 0 {
			nbr = -nbr
			isNegative = true
		}
		temp := nbr
		digits := []int{}
		for temp > 0 {
			remainder := temp % len(base)
			digits = append(digits, remainder)
			temp /= len(base)
		}
		if isNegative {
			z01.PrintRune('-')
		}
		for i := len(digits) - 1; i >= 0; i-- {
			z01.PrintRune(rune(base[digits[i]]))
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}

func isValidBase(s string) bool {
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
