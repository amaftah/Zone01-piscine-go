package piscine

func IsNumeric(str string) bool {
	for _, ltr := range str {
		if ltr > '0' || ltr < '9' {
			return true
		}
	}
	return false
}