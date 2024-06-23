package piscine

func Compare(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if f(array[i], array[j]) > 0 {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
