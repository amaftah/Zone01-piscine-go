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
    } else if len(arguments) == 0 {
        return
    }

    for _, arg := range arguments {
        n := 0
        valid := true
        for _, c := range arg {
            if c < '0' || c > '9' {
                valid = false
                break
            }
            n = n*10 + int(c-'0')
        }
        if !valid || n < 1 || n > 26 {
            z01.PrintRune(' ')
            continue
        }
        if upper {
            z01.PrintRune(rune('A' + n - 1))
        } else {
            z01.PrintRune(rune('a' + n - 1))
        }
    }
    z01.PrintRune('\n')
}