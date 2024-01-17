package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := func(nm int) int {
		cnt := 0
		for nm != 0 {
			nm /= 10
			cnt++
		}
		return cnt
	}
	nmdigits := digits(n)

	dgts := make([]int, nmdigits)
	for i := nmdigits - 1; i >= 0; i-- {
		dgts[i] = n % 10
		n /= 10
	}

	for i := 0; i < nmdigits-1; i++ {
		for j := 0; j < nmdigits-i-1; j++ {
			if dgts[j] > dgts[j+1] {
				dgts[j], dgts[j+1] = dgts[j+1], dgts[j]
			}
		}
	}

	for _, dgt := range dgts {
		z01.PrintRune(rune(dgt) + '0')
	}
}
