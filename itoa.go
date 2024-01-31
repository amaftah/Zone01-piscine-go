package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result string
	isNegative := n < 0
	if isNegative {
		n = -n
	}

	for n > 0 {
		result = string('0'+n%10) + result
		n = n / 10
	}

	if isNegative {
		result = "-" + result
	}

	return result
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
