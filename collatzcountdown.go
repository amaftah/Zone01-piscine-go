package main

import "fmt"

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	count := 0
	for start != 1 {
		if start%2 == 0 {
			start /= 2
		} else {
			start = start*3 + 1
		}
		count++
	}
	return count
}

func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}
