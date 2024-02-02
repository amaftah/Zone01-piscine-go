package main

import "fmt"

func FishAndChips(n int) string {
	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	}
	if n%2 == 0 {
		return "fish"
	}
	if n%3 == 0 {
		return "chips"
	}
	if n != n%2 && n != n%3 {
		return "error: non divisible"
	}
	if n%2 < 0 && n%3 < 0 {
		return "error: number is negative"
	}
	return ""
}

func main() {
	fmt.Println(FishAndChips(-4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}
