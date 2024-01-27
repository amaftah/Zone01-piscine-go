package main

import "github.com/01-edu/z01"

func main() {
	QuadE(5, 1)
}

func QuadE(x int, y int) string {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				switch {
				case (i == 0 && j == 0):
					z01.PrintRune('A')
				case (i == 0 && j == 0) || (i == y-1 && j == x-1):
					if j == x-1 && i == 0 {
						z01.PrintRune('C')
					} else if j == 0 && i == y-1 {
						z01.PrintRune('C')
					} else {
						z01.PrintRune('A')
					}
				case (i == 0 && j == x-1) || (i == y-1 && j == 0):
					z01.PrintRune('C')
				case i == 0 || i == y-1:
					z01.PrintRune('B')
				case j == 0 || j == x-1:
					z01.PrintRune('B')
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
