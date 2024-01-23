package piscine

func Join(strs []string, sep string) string {
	var def string
	for i, word := range strs {
		if i == 0 {
			def += word
		} else {
			def += sep + word
		}
	}
	return def
}