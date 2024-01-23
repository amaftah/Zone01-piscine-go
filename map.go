package piscine

func Map(f func(int) bool, a []int) []bool {
    result := make([]bool, len(a))
    for i, v := range a {
        result[i] = f(v)
    }
    return result
}