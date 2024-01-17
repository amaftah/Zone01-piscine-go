package piscine

func TrimAtoi(s string) int {
	var res int
	var sign int = 1
	for _, c := range s {
		if c >= '0' && c <= '9' {
			res = res*10 + int(c-'0')
		} else if c == '-' && res == 0 {
			sign = -1
		}
	}
	return res * sign
}