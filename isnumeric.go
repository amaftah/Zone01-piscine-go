package main

import "fmt"

func IsNumeric(str string) bool {
	for _, ltr := range str {
		if ltr < '0' || ltr > '9' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsNumeric("Hello!"))
	fmt.Println(IsNumeric("Salut!"))
	fmt.Println(IsNumeric("Ola!"))
}