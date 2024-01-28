package main

import (
	"fmt"
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

	if len(matches) == 0 {
		return "Not a quad function\n"
	}

	sort.Strings(matches)
	return strings.Join(matches, " || ") + "\n"
}

func main() {
	fmt.Println(quadchecker(1, 1))
}
