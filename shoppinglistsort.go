package main

import "fmt"

func ShoppingListSort(slice []string) []string {
	for s := 0; s < len(slice)-1; s++ {
		for i := 0; i < len(slice)-1; i++ {
			if len(slice[i]) > len(slice[i+1]) {
				slice[i+1], slice[i] = slice[i], slice[i+1]
			}
		}
	}
	return slice
}

func main() {
	slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
	fmt.Println(ShoppingListSort(slice))
}