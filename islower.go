package piscine

func IsLower(str string) bool {
	for _, ltr := range str {
		if ltr > 'a' || ltr < 'z' {
			return true
		}
	}
	return false
}