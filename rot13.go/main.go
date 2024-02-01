package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	var output string

	for _, r := range input {
		switch {
		case 'a' <= r && r <= 'z':
			output += string((r-'a'+13)%26 + 'a')
		case 'A' <= r && r <= 'Z':
			output += string((r-'A'+13)%26 + 'A')
		default:
			output += string(r)
		}
	}
	fmt.Println(output)
}
