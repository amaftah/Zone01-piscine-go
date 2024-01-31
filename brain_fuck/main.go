package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <brainfuck code>")
		os.Exit(1)
	}

	code := os.Args[1]
	memory := make([]byte, 2048)
	ptr := 0
	loopStack := make([]int, 0)

	for pc := 0; pc < len(code); pc++ {
		switch code[pc] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '.':
			fmt.Print(string(memory[ptr]))
		case '[':
			loopStack = append(loopStack, pc)
		case ']':
			if memory[ptr] != 0 {
				pc = loopStack[len(loopStack)-1]
			} else {
				loopStack = loopStack[:len(loopStack)-1]
			}
		}
	}
}
