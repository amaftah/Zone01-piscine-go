package main

import (
	"fmt"
	"os"
)

func alphaMirror(s string) {
	result := ""
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			if r >= 'A' && r <= 'Z' {
				result += string('Z' - (r - 'A'))
			} else {
				result += string('z' - (r - 'a'))
			}
		} else {
			result += string(r)
		}
	}
	fmt.Println(result)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	alphaMirror(os.Args[1])
}
