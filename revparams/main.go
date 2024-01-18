package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := len(arg) - 1; i > 0; i-- {
		for _, letter := range arg[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}