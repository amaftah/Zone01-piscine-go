package main

import "fmt"

func AlphaCount(s string) int {
	cnt := 0
	for _, letter := range s {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			cnt++
		}
	}
	return cnt
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}