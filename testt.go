package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountChar("Hello World", 'l'))                  // 'l' appears 3 times
	fmt.Println(CountChar("5  balloons", 5))                    // '5' appears 0 times
	fmt.Println(CountChar("        yufyu     gfrsdgrdhd", ' ')) // ' ' appears 3 times
	fmt.Println(CountChar(" ", ' '))                            // '7' appears 1 time
}

func CountChar(str string, c rune) int {
	count := 0

	for _, char := range str {
		if char == c && c != ' ' {
			count++
		} else if c == ' ' && c == char {
			count++
			break
		}
	}

	return count
}
