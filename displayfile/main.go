package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("File name missing")
        return
    }

    if len(os.Args) > 2 {
        fmt.Println("Too many arguments")
        return
    }

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error opening file: ", err)
        return
    }
	defer file.Close()

    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println("Error reading file: ", err)
        return
    }

    fmt.Print(string(data))
}