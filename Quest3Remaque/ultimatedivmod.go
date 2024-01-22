package piscine

func UltimateDivMod(a *int, b *int) {
	tempdiv := *a / *b
	tempmod := *a % *b
	*a = tempdiv
	*b = tempmod
}
