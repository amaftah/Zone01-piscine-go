package main

import (
    "fmt"
)

type Food struct {
    preptime int
}

var menu = map[string]Food{
    "burger":  {15},
    "chips":   {10},
    "nuggets": {12},
}

func FoodDeliveryTime(order string) int {
    if food, ok := menu[order]; ok {
        return food.preptime
    } else {
        return 404
    }
}

func main() {
    fmt.Println(FoodDeliveryTime("burger"))
    fmt.Println(FoodDeliveryTime("chips"))
    fmt.Println(FoodDeliveryTime("pizza"))
    fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))
}