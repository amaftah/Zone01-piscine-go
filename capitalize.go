package piscine

func Capitalize(s string) string {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] >= 'A' && slice[i] <= 'Z' {
			if i != 0 {
				if slice[i-1] >= 'a' && slice[i-1] <= 'z' {
					continue
				} else if slice[i-1] >= 'A' && slice[i-1] <= 'Z' {
					continue
				} else if slice[i-1] >= '0' && slice[i-1] <= '9' {
					continue
				} else {
					slice[i] = slice[i] + 32
				}
			}
		} else if slice[i] >= 'a' && slice[i] <= 'z' {
			if i == 0 {
				slice[i] = slice[i] - 32
			} else if slice[i-1] >= 'a' && slice[i-1] <= 'z' {
				continue
			} else if slice[i-1] >= 'A' && slice[i-1] <= 'Z' {
				continue
			} else if slice[i-1] >= '0' && slice[i-1] <= '9' {
				continue
			} else {
				slice[i] = slice[i] - 32
			}
		}
	}
	return string(slice)
}