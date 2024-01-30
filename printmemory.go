package main

import "fmt"

func PrintMemory(arr [10]byte) {
	for i := 0; i < len(arr); i += 16 {
		for j := i; j < i+16 && j < len(arr); j++ {
			fmt.Printf("%02X ", arr[j])
		}
		fmt.Print(" ")
		for j := i; j < i+16 && j < len(arr); j++ {
			if arr[j] >= 32 && arr[j] <= 126 {
				fmt.Printf("%c", arr[j])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
