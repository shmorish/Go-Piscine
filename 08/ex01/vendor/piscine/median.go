package piscine

func IntVecLen(s []int) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func Median(arr ...int) int {
	for i := 0; i < IntVecLen(arr); i++ {
		for j := i; j < IntVecLen(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr[IntVecLen(arr)/2]
}
