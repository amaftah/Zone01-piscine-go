package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func quadchecker(input string) string {
	quads := []string{"./QuadA", "./QuadB", "./QuadC", "./QuadD", "./QuadE"}
	matches := []string{}

	for _, quad := range quads {
		cmd := exec.Command(quad)
		out, _ := cmd.Output()

		if strings.TrimSpace(string(out)) == input {
			matches = append(matches, strings.TrimPrefix(quad, "./"))
		}
	}

	ex1 := exec.Command("./QuadA")
	ex1.Stdout = os.Stdout
	ex1.Stderr = os.Stderr
	err := ex1.Run()
	if err != nil {
		fmt.Println("Error running QuadA:", err)
	}

	ex2 := exec.Command("./QuadB")
	ex2.Stdout = os.Stdout
	ex2.Stderr = os.Stderr
	err = ex2.Run()
	if err != nil {
		fmt.Println("Error running QuadB:", err)
	}

	ex3 := exec.Command("./QuadC")
	ex3.Stdout = os.Stdout
	err = ex3.Run()
	if err != nil {
		fmt.Println("Error running QuadC:", err)
	}

	if len(matches) == 0 {
		return "Not a quad function\n"
	}

	sort.Strings(matches)
	return strings.Join(matches, " || ") + "\n"
}
