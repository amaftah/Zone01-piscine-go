package main

import (
    "os"
)

func Atoi(s string) (int, bool) {
    isNegative := false
    if s != "" && s[0] == '-' {
        isNegative = true
        s = s[1:]
    }

    result := 0
    for _, c := range s {
        if c < '0' || c > '9' {
            return 0, false
        }
        result = result*10 + int(c-'0')
    }

    if isNegative {
        result = -result
    }

    return result, true
}

func WriteInt(n int) {
    if n < 0 {
        os.Stdout.WriteString("-")
        n = -n
    }

    if n/10 != 0 {
        WriteInt(n / 10)
    }

    os.Stdout.WriteString(string('0' + n%10))
}

func main() {
    if len(os.Args) != 4 {
        return
    }

    a, ok1 := Atoi(os.Args[1])
    operator := os.Args[2]
    b, ok2 := Atoi(os.Args[3])

    if !ok1 || !ok2 {
        return
    }

    switch operator {
    case "+":
        WriteInt(a + b)
    case "-":
        WriteInt(a - b)
    case "*":
        WriteInt(a * b)
    case "/":
        if b == 0 {
            os.Stdout.WriteString("No division by 0")
        } else {
            WriteInt(a / b)
        }
    case "%":
        if b == 0 {
            os.Stdout.WriteString("No modulo by 0")
        } else {
            WriteInt(a % b)
        }
    default:
        return
    }
    os.Stdout.WriteString("\n")
}