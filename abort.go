package main

import "fmt"

func Abort(a, b, c, d, e int) int {
	var arr [5]int
	arr[0] = a
	arr[1] = b
	arr[2] = c
	arr[3] = d
	arr[4] = e
	for i := 0; i < 5; i++ {
		for j := 0; j < 4-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}
	}
	return arr[2]
}

func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
