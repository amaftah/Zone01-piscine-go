package main

import "fmt"

func LastRune(s string) rune {
	r := []rune(s)
	return r[len(r)-1]
}

func main() {

	fmt.Println(LastRune("Hello!"))
	fmt.Println(LastRune("Salut!"))
	fmt.Println(LastRune("Ola!"))
}
