package main

import (
	"fmt"
	
)

func BasicAtoi(s string) int {
	var result int
	for _, v := range s {
		result = result * 10 + int(v-'0')
	}
	return result
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
