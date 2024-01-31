package main

import "fmt"

func BinaryCheck(nbr int) int {
	if nbr > 0 {
		return BinaryCheck(-nbr)
	}
	if nbr == 0 {
		return 1
	}
	if nbr < 0 {
		return 0
	}
	return nbr
}

func main() {
	fmt.Println(BinaryCheck(5))
	fmt.Println(BinaryCheck(0))
	fmt.Println(BinaryCheck(8))
	fmt.Println(BinaryCheck(-9))
	fmt.Println(BinaryCheck(-4))
}
