package piscine

func isAlphanumeric(c rune) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func Capitalize(s string) string {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if isAlphanumeric(slice[i]) {
			if i == 0 || (i > 0 && !isAlphanumeric(slice[i-1])) {
				if slice[i] >= 'a' && slice[i] <= 'z' {
					slice[i] -= 32
				}
			} else if slice[i] >= 'A' && slice[i] <= 'Z' {
				slice[i] += 32
			}
		}
	}
	return string(slice)
}