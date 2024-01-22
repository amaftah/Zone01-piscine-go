package main

import (
	"github.com/01-edu/z01"

	"os"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	}
	return false
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	arg := os.Args[1:]
	if isEven(len(arg)) {
		printStr(EvenMsg + "\n")
	} else {
		printStr(OddMsg + "\n")
	}
}