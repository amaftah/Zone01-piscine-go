package main

import "github.com/01-edu/z01"

func main() {
	digits := '0'
	for digits <= '9' {
		z01.PrintRune(digits)
		digits++
	}
	z01.PrintRune('\n')


}