package piscine

func StrLen(str string) int {
	var count int
	for range str {
		count++
	}
	return count
}

func NRune(s string, n int) rune {
	if n < 0 {
		return 0
	}
	runes := []rune(s)
	if n > StrLen(s) {
		return 0
	}
	return runes[n-1]
}
