package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	length := max - min
	answer := make([]int, length)
	for i := 0; i < length; i++ {
		answer[i] = min + i
	}
	return answer
}
