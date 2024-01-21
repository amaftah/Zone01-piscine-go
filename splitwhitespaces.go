package main

import	"fmt"

func SplitWhiteSpaces(s string) []string {
	a := make([]string, 0)
	b := make([]byte, 0)
	for _, i := range s {
		if i == ' ' || i == '\t' || i == '\n' {
			if len(b) != 0 {
				a = append(a, string(b))
				b = make([]byte, 0)
			}
		} else {
			b = append(b, byte(i))
		}
	}
	if len(b) != 0 {
		a = append(a, string(b))
	}
	return a
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
