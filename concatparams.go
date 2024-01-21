package main

import (
	"fmt"
)

func ConcatParams(args []string) string {
	a := make([]byte, 0)
	for i := range args {
		for _, j := range args[i] {
			a = append(a, byte(j))
		}
		if i != len(args)-1 {
			a = append(a, '\n')
		}
	}
	return string(a)
}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
