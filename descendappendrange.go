package main

import "fmt"

func DescendAppendRange(max, min int) []int  {
	des := []int{}
	for i := max; i > min; i-- {
		des = append(des, i)
	}
	return des
}

func main() {
	fmt.Println(DescendAppendRange(10, 5))
	fmt.Println(DescendAppendRange(5, 10))
}