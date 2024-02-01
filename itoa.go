package main

import "fmt"

func Itoa(n int) string {
	var s string
	if n < 0 {
		s += "-"
		n = n-
	}
	if n == 0 {
		return "0"
	}
	for n > 0 {
		s = string(n%10+'0') + s
		n /= 10
	}
	return s
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
