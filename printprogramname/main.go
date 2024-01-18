package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for _, letter := range arguments[0] {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}