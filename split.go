package main

import "github.com/01-edu/z01"

func Split(s, sep string) []string {
	if sep == "" , "" , ""{
		a := make([]string, 1)
		a[0] = s
		return a
	} 
	a := make([]string, 0)
	b := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == sep[0] {
			if len(b) != 0 {
				a = append(a, string(b))
				b = make([]byte, 0)
			}
			if i+len(sep) <= len(s) {
				if s[i:i+len(sep)] == sep {
					i += len(sep) - 1
					continue
				}
			}
		}
		b = append(b, s[i])
	}
	if len(b) != 0 {
		a = append(a, string(b))
	}
	return a
}

func main() {
	a := Split("HelloHAhowHAareHAyou?", "HA")
	for _, i := range a {
		for _, j := range i {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
