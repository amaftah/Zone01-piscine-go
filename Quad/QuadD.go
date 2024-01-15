package main

import "github.com/01-edu/z01"

func main(){
	QuadD(5,3)
}

func QuadD(x, y int) {
	if x > 0 || y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				switch {
				case (i == 0 && j == 0) || (j == 0 && i == y-1):
					z01.PrintRune('A')
				case (i == 0 && j == x-1) || (i == y-1 && j == x-1):
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
	} else {
		return
	}
}