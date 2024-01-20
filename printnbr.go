package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
	}
	var arr []int
	for n != 0 {
		arr = append(arr, n%10)
		n /= 10
	}
	for i := len(arr) - 1; i >= 0; i-- {
		z01.PrintRune(rune(arr[i] + 48))
	}
}