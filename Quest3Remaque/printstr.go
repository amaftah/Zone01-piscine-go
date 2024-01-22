package main

import "github.com/01-edu/z01"

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}