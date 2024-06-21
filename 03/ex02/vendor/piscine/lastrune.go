package piscine

func StrLen(str string) int {
	var count int
	for range str {
		count++
	}
	return count
}

func LastRune(s string) rune {
	runes := []rune(s)
	return runes[StrLen(s)-1]
}
