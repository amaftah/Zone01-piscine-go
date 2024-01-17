package main

import "fmt"

func IsPrintable(str string) bool {
	for _, ltr := range str {
		if ltr > 31 || ltr < 127 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(IsPrintable("Hello!"))
}