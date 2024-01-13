package main

import "github.com/01-edu/z01"

func ConvertIntToString(num int) string {
	if num == 0 {
		return "0"
	} else if num == 1 {
		return "1"
	} else if num == 2 {
		return "2"
	} else if num == 3 {
		return "3"
	} else if num == 4 {
		return "4"
	} else if num == 5 {
		return "5"
	} else if num == 6 {
		return "6"
	} else if num == 7 {
		return "7"
	} else if num == 8 {
		return "8"
	} else {
		return "9"
	}
}

func IntToAlpha(num int) string {
	var result string

	if num == 0 {
		return "0"
	}
	for num > 0 {
		result = ConvertIntToString(num%10) + result
		num /= 10
	}
	return result
}

func PrintString(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
}

func PrintCombinations(n int, prev int, result string, count *int) {
	for i := 0; i < 10; i++ {
		if prev < i {
			if n == 1 {
				if *count > 0 {
					PrintString(", ")
				}
				PrintString(result + IntToAlpha(i))
				*count++
			} else {
				PrintCombinations(n-1, i, result+IntToAlpha(i), count)
			}
		}
	}
}

func PrintCombinationsN(n int) {
	var count int = 0
	for i := 0; i < 10; i++ {
		if n > 1 {
			PrintCombinations(n-1, i, IntToAlpha(i), &count)
		} else {
			if count > 0 {
				PrintString(", ")
			}
			PrintString(IntToAlpha(i))
			count++
		}
	}
	PrintString("\n")
}

func main() {
	n := 3
	
	PrintCombinationsN(n)
}
