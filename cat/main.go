package main

import (
    "github.com/01-edu/z01"
    "io/ioutil"
    "os"
)

func printError(err error) {
    errorMsg := "ERROR: " + err.Error()
    for _, r := range errorMsg {
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
}

func main() {
    if len(os.Args) == 1 {
        data, err := ioutil.ReadAll(os.Stdin)
        if err != nil {
            printError(err)
            os.Exit(1)
        }
        os.Stdout.Write(data)
    } else {
        for _, filename := range os.Args[1:] {
            file, err := os.Open(filename)
            if err != nil {
                printError(err)
                os.Exit(1)
            }
            defer file.Close()

            data, err := ioutil.ReadAll(file)
            if err != nil {
                printError(err)
                os.Exit(1)
            }

            os.Stdout.Write(data)
        }
    }
}