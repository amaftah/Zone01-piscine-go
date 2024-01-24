package piscine

func Unmatch(a []int) int {
	CntMatch := make(map[int][int])

	for _, v := range a {
		CntMatch[v]++
	}

	for _, num := range CntMatch {
		if num%2 == 1 {
			return num
		}
	}
	return -1
}