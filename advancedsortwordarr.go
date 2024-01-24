package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	x := len(a)
	for i := 0; i < x; i++ {
		for j := i + 1; j < x; j++ {
			if f(a[i], a[j]) > 0 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}