package piscine

func SortWordArr(a []string) {
	x := len(a)
	for i := 0; i < x; i++ {
		for j := i + 1; j < x; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}