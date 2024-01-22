package main

import "github.com/01-edu/z01"

func main() {
	ReverseAlpha := 'z'
	for ReverseAlpha >= 'a' {
		z01.PrintRune(ReverseAlpha)
		ReverseAlpha--
	}
	z01.PrintRune('\n')


}