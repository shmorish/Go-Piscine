package piscine

func IntVecLen(s []int) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func IsSorted(f func(int, int) int, a []int) bool {
	if IntVecLen(a) <= 1 {
		return true
	}
	for i := 1; i < IntVecLen(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			return false
		}
	}
	return true
}
