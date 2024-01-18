package main

import ( "os"
		"github.com/01-edu/z01"	)

func main() {
	params := os.Args
	for i := 1; i < len(params); i++ {
		for _, letter := range params[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}