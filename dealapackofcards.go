package main

import "fmt"

func DealAPackOfCards (deck []int) {
	fmt.Printf("Player1 %v, %v, %v\n", deck[0], deck[1], deck[2])
	fmt.Printf("Player2 %v, %v, %v\n", deck[3], deck[4], deck[5])
	fmt.Printf("Player3 %v, %v, %v\n", deck[6], deck[7], deck[8])
	fmt.Printf("Player4 %v, %v, %v\n", deck[9], deck[10], deck[11])
}


func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}