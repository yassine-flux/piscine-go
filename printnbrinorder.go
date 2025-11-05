package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}

	digits := [10]int{}
	temp := n
	for temp > 0 {
		digit := temp % 10
		digits[digit]++
		temp /= 10
	}

	for i := 0; i <= 9; i++ {
		for j := 0; j < digits[i]; j++ {
			z01.PrintRune(rune(i + '0'))
		}
	}
}
