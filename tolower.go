package piscine

func ToLower(s string) string {
	var rslt string
	for _, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			letter += 32
		}
		rslt += string(letter)
	}
	return rslt
}