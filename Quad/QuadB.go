package main

import "github.com/01-edu/z01"

func QuadB(x int, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				switch {
				case (i == 0 && j == 0):
					z01.PrintRune('/')
				case (i == 0 && j == 0) || (i == y-1 && j == x-1):
					if j == x-1 && i == 0 {
						z01.PrintRune('\\')
					} else if j == 0 && i == y-1 {
						z01.PrintRune('\\')
					} else {
						z01.PrintRune('/')
					}
				case (i == 0 && j == x-1) || (i == y-1 && j == 0):

					z01.PrintRune('\\')
				case i == 0 || i == y-1:
					z01.PrintRune('*')
				case j == 0 || j == x-1:
					z01.PrintRune('*')
				default:
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}

func main() {
	QuadB(1,5)
}