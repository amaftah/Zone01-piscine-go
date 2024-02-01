package main

import "fmt"

/*func atoi(s string) int {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	if i == len(s) {
		return 0
	}

	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	num := 0
	for ; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			break
		}
		num = num*10 + int(ch-'0')
	}

	return sign * num
}*/

func itoa(n int) string {
	s := ""
	if n < 0 {
		s += "-"
		n = -n
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
