package piscine

func IsUpper(str string) bool {
	for _, ltr := range str {
		if ltr > 'A' || ltr < 'Z' {
			return true
		}
	}
	return false
}