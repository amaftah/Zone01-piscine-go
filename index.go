package main

import "fmt"

func Index(s string, toFind string) int {
	for r := 0; r < len(s)-len(toFind); r++ {
		if s[r:r+len(toFind)] == toFind{
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