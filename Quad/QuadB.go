package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil {
		return
	}

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
