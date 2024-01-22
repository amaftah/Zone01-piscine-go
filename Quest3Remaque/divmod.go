package main
import "github.com/01-edu/z01"

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}