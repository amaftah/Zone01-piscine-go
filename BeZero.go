package main

import "fmt"

func BeZero(slice []int) []int {
	if len(slice) == 0 {
		var empty []int
		return empty
	}

	for i := 0; i < len(slice); i++ {
		slice[i] = 0
	}
	return slice
}

func main() {
	fmt.Println(BeZero([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(BeZero([]int{}))
}
