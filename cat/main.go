package main

import (
    "github.com/01-edu/z01"
    "io/ioutil"
    "os"
)

func main() {
    if len(os.Args) == 1 {    
        data, err := ioutil.ReadAll(os.Stdin)
        if err != nil {
            z01.PrintRune('\n')
            return
        }
        os.Stdout.Write(data)
    } else {
        file, err := os.Open(os.Args[1])
        if err != nil {
            z01.PrintRune('\n')
            return
        }
        defer file.Close()

        data, err := ioutil.ReadAll(file)
        if err != nil {
            z01.PrintRune('\n')
            return
        }

        os.Stdout.Write(data)
    }
}