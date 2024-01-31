package main

import "fmt"

func AtoiBase(s string, base string) int {
	var nbr int

	if len(base) < 2 {
		return 0
	}

	sign := 1

	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}

	for _, c := range s {
		nbr *= len(base)
		for i, v := range base {
			if c == v {
				nbr += i
			}
			if v == '-' || v == '+' {
				return 0
			}
		}
	}
	return nbr * sign
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
