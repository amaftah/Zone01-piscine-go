package main

import "github.com/01-edu/z01"

func QuadA(x, y int) string {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				switch {
				case (i == 0 || i == y-1) && (j == 0 || j == x-1):
					z01.PrintRune('o')
				case i == 0 || i == y-1:
					z01.PrintRune('-')
				case j == 0 || j == x-1:
					z01.PrintRune('|')
				default:
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
	return itoa(x) + itoa(y)
}

func itoa(n int) string {
	var res string
	for n != 0 {
		res = string(n%10+48) + res
		n /= 10
	}
	return res
}

func main() {
	QuadA(5, 3)
}
