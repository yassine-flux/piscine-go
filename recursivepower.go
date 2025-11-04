package piscine

func RecursivePower(nb, power int) int {
	result := 1
	if power < 0 {
		return 0
	}
	if power == 1 {
		return nb
	}
	if power == 0 {
		return 1
	}
	result = nb * RecursivePower(nb, power-1)
	return result
}
