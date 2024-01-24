package main

import "fmt"


func ShoppingSummaryCounter(str string) map[string]int {
	a := 1
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			a++
		}
	}
	arr := make([]string, a)
	b := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			b++
		} else {
			arr[b] += string(str[i])
		}
	}
	m := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	return m 
}
func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}
