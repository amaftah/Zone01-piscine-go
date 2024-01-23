package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
    if len(a) <= 1 {
        return true
    }

    direction := f(a[0], a[1])

    for i := 1; i < len(a)-1; i++ {
        if direction > 0 && f(a[i], a[i+1]) < 0 {
            return false
        } else if direction < 0 && f(a[i], a[i+1]) > 0 {
            return false
        } else if direction == 0 {
            direction = f(a[i], a[i+1])
        }
    }

    return true
}