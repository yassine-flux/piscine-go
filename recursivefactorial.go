package piscine

func RecursiveFactorial(nb int) int {
	result := 1
	if nb >= 21 || nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	} else {
		result = nb * RecursiveFactorial(nb-1)
	}
	return result
}
