package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb >= 21 || nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	for i := nb; i >= 1; i-- {
		result *= i
	}
	return result
}
