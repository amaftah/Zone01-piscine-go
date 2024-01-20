package main

import (
	"github.com/01-edu/z01"

	"os"
)

func main() {

	args := os.Args[1:]
	len := 0
	for i := range args {
		len = i + 1
	}
	for i := 0; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}
	for i := range args {
		for _, c := range args[i] {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}