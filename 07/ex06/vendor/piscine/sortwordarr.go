package piscine

func VecLen(s []string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func SortWordArr(array []string) {
	for i := 0; i < VecLen(array); i++ {
		for j := i + 1; j < VecLen(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
