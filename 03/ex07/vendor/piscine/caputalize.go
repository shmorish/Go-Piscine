package piscine

func StrLen(str string) int {
	var count int
	for range str {
		count++
	}
	return count
}

func IsAlnum(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return true
	}
	return false
}

func Capitalize(word string) string {
	runes := []rune(word)
	if runes[0] >= 'a' && runes[0] <= 'z' {
		runes[0] = runes[0] - 32
	}
	for i := 1; i < StrLen(word); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			if IsAlnum(runes[i-1]) {
				runes[i] = runes[i] + 32
			}
		} else if runes[i] >= 'a' && runes[i] <= 'z' {
			if !IsAlnum(runes[i-1]) {
				runes[i] = runes[i] - 32
			}
		}
	}
	return string(runes)
}
