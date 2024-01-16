package main

import "fmt"

func Index(s string, toFind string) int {
	for r := 0; r < len(s); r++ {
		if s[r] == toFind[0] {
			return r
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}