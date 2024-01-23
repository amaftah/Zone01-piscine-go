package piscine

func BasicJoin(elems []string) string {
	var def string
	for _, word := range elems {
		def += word
	}
	return def
}
