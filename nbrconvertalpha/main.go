package main

import (
    "os"
    
    "github.com/01-edu/z01"
)

func main() {
    arguments := os.Args[1:]
    upper := false
    if len(arguments) > 0 && arguments[0] == "--upper" {
        upper = true
        arguments = arguments[1:]
    }


    convert := func(s string) rune {
        n := 0
        for _, c := range s {
            if c < '0' || c > '9' {
                return ' '
            }
            n = n*10 + int(c-'0')
        }
        if n < 1 || n > 26 {
            return ' '
        }
        if upper {
            return rune('A' + n - 1)
        }
        return rune('a' + n - 1)
    } 

    for _, word := range arguments {
        letter := convert(word)
        z01.PrintRune(letter)
    }
    z01.PrintRune('\n')
}