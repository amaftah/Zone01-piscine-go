package main

import (
	
	"github.com/01-edu/z01"

	"piscine"
)

func IsNegative(nb int) {
	for nb < 0 {
		z01.PrintRune('T')
	} else { 
		z01.PrintRune('F')
	}
}

func main() {
	piscine.IsNegative(1)
	piscine.IsNegative(0)
	piscine.IsNegative(-1)
}
