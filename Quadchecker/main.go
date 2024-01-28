package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type QuadFunc func(int, int) int {
	return x * y
}

func quadChecker(x int, y int, f QuadFunc) {
	if x > 0 && y > 0 {
		f(x, y)
	} else {
		fmt.Println("Error: Both width and height must be positive integers.")
	}
}

func main() {
	// Create a temporary file to hold the Quadchecker code.
	tmpfile, err := ioutil.TempFile("", "quadchecker")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(tmpfile.Name())

	// Write the Quadchecker code to the temporary file.
	code := `
}

func QuadA(x int, y int) {
	fmt.Printf("QuadA: %d\n", x*y)
}
func QuadB(x int, y int) {
	fmt.Printf("QuadB: %d\n", x*y)
}
func QuadC(x int, y int) {
	fmt.Printf("QuadC: %d\n", x*y)
}
func QuadD(x int, y int) {
	fmt.Printf("QuadD: %d\n", x*y)
}

func QuadE(x int, y int) {
	fmt.Printf("QuadE: %d\n", x*y)
}
`
	tmpfile.Write([]byte(code))

	// Run the Quadchecker code.
	cmd := exec.Command("go", "run", tmpfile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	// Call the Quadchecker function.
	quadChecker(5, 10, QuadA)
	quadChecker(5, 10, QuadB)
	quadChecker(5, 10, QuadC)
	quadChecker(5, 10, QuadD)
	quadChecker(5, 10, QuadE)
}
