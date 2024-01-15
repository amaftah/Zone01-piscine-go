package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 0
	} else {
		result := 1
		for x := 1; x <= power; x++ {
			result *= nb
		}
		return result
	}
}
