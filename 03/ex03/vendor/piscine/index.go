package piscine

func StrLen(str string) int {
	var count int
	for range str {
		count++
	}
	return count
}

func Index(s string, toFind string) int {
	for i := 0; i < StrLen(s); i++ {
		if s[i] == toFind[0] {
			if toFind == s[i:i+StrLen(toFind)] {
				return i
			}
		}
	}
	return -1
}
