package main

import (
    "os"
)

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
            return
        }
        WriteInt(a / b)
    case "%":
        if b == 0 {
            return
        }
        WriteInt(a % b)
    default:
        return
    }
}

func Atoi(s string) (int, bool) {
    if len(s) == 0 {
        return 0, false
    }

    result := 0
    for _, r := range s {
        if r < '0' || r > '9' {
            return 0, false
        }
        result = result*10 + int(r-'0')
    }

    return result, true
}

func WriteInt(n int) {
    buf := make([]byte, 20)
    i := len(buf)
    isNegative := false

    if n == 0 {
        i--
        buf[i] = '0'
    } else {
        if n < 0 {
            isNegative = true
            n = -n
        }
        for n > 0 {
            i--
            buf[i] = byte(n%10) + '0'
            n /= 10
        }
        if isNegative {
            i--
            buf[i] = '-'
        }
    }
    os.Stdout.Write(buf[i:])
    os.Stdout.Write([]byte{'\n'})
}