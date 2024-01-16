package Piscine

func IsPrintable(str string) bool {
	for _, ltr := range str {
		if ltr > 31 || ltr < 127 {
			return true
		}
	}
	return false
}