package main

import "fmt"

func IsNumeric(str string) bool {
	for _, ltr := range str {
		if ltr < '0' || ltr <{	
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