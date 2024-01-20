package main

import "fmt"

func NRune(s string, n int) rune {
	//turn string into rune slice
	var res rune
	for i, v := range s {
		if i == n-1 {
			res = v
		}
	}
	return res
}

func main() {
	fmt.Println(NRune("Hello!", 3))
	fmt.Println(NRune("Salut!", 2))
	fmt.Println(NRune("Ola!", 4))
	fmt.Println(NRune("Bye!", 1))
	fmt.Println(NRune("Salut!", 0))
}