package main

import "fmt"

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return NbrBase(AtoiBase(nbr, baseFrom), baseTo)
}

func NbrBase(nbr int, base string) string {
	var result string
	var n int
	for nbr != 0 {
		n = nbr % len(base)
		nbr = nbr / len(base)
		result = string(base[n]) + result
	}
	return result
}

func AtoiBase(s string, base string) int {
	var result int
	for _, r := range s {
		for i, b := range base {
			if r == b {
				result = result*len(base) + i
			}
		}
	}
	return result
}

func main() {
	result := ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}