package main

import "github.com/01-edu/z01"

func PrintComb() {
	for a := '0'; a <= '9'; a++ {
		for b := '1'; a <= '9'; b++ {
			for c := '2'; c <= '9'; c++ {
				if a < b && b < c {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)

					if a == '7' && b == '8' && c == '9' {
						z01.PrintRune(10)
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
func main() {
	PrintComb()
}
