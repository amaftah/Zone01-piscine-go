package main

import "fmt"

func Compact(ptr *[]string) int {
	original := *ptr
	count := 0


	for _, item := range original {
		if item != "" {
			count++
		}
	}

	compacted := make([]string, count)

	
	






const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}
