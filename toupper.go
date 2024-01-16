package main

func ToUpper(s string) string {
	var rslt string
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			letter -= 32
		}
		rslt += string(letter)
	}
	return rslt
}