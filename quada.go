package main

import "fmt"

func main() {
	QuadA(50,30)
}

func QuadA(x,y int) {
	if x > 0 || y > 0 {
		for a := 0 ; a < y; a++ {
			for b := 0; b < x; b++ {
			if (a == 0 || a == y-1) || (b == 0 || b == x-1) {
					if a == 0 || a == y-1 {
							if b == 0 || b == x-1 {
								fmt.Printf("o")
							} else {
								fmt.Printf("-")
						}
					} else {
						fmt.Printf("|")
					}

			} else {
				fmt.Printf(" ")
				}	
			}
			fmt.Println()
		}
	} else {
		 return
	}
}
