package main

import "github.com/01-edu/z01"



func Rot14(str string) string {
    slice := []rune(str)
    for i := 0; i < len(slice); i++ {
        if slice[i] >= 'A' && slice[i] <= 'Z' {
            slice[i] = 'A' + (slice[i]-'A'+14)%26
        } else if slice[i] >= 'a' && slice[i] <= 'z' {
            slice[i] = 'a' + (slice[i]-'a'+14)%26
        }
    }
    return string(slice)
}
func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
