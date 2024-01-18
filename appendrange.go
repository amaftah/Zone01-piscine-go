package main

import "fmt"

func AppendRange(min, max int) []int {
	var rslt []int
	if min >= max {
		return rslt
	}
	for s := min; s < max; s++ {
		rslt = append(rslt, s)
	}
	return rslt
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}