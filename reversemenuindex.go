package piscine

func ReverseMenuIndex(menui []string) []string {
	result := make([]string, len(menu))
	for a, b := 0, len(menu)-1; a < len(result); a, b = a+1, b-1 {
		result[a] = menu[b]
	}
	return result
}