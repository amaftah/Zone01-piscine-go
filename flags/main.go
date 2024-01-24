package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
        fmt.Println("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.")
        fmt.Println("--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
        return
    }

    var insert string
    var order bool
    var arg string

    for i := 1; i < len(os.Args); i++ {
        if len(os.Args[i]) > 9 && os.Args[i][:9] == "--insert=" {
            insert = os.Args[i][9:]
        } else if len(os.Args[i]) > 3 && os.Args[i][:3] == "-i=" {
            insert = os.Args[i][3:]
        } else if os.Args[i] == "--order" || os.Args[i] == "-o" {
            order = true
        } else {
            arg = os.Args[i]
        }
    }

    if !order {
        if insert != "" {
            fmt.Println(arg + insert)
        } else {
            fmt.Println(arg)
        }
    }

    if order {
        fmt.Println(Order(arg + insert))
    }
}

func Order(s string) string {
    r := []rune(s)
    for i := 0; i < len(r); i++ {
        for j := i + 1; j < len(r); j++ {
            if r[j] < r[i] {
                r[i], r[j] = r[j], r[i]
            }
        }
    }
    return string(r)
}