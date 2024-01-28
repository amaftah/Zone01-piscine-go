package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func checkDimensions(filename string) {
	// Read the file
	data, err := ioutil.ReadFile("./QuadA")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	// Here you can add your logic to check the dimensions of the file
	// For now, we just print the length of the data
	fmt.Println("Length of file:", len(data))
}

func main() {
	// Check if a filename was provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	// Loop over all provided filenames
	for _, filename := range os.Args[1:] {
		checkDimensions(filename)
	}
}
