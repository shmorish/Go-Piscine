package piscine

func Atoi(s string) int {
	var result int
	var sign int = 1
	for i, c := range s {
		if i == 0 && c == '-' {
			sign = -1
			continue
		}
		if i == 0 && c == '+' {
			continue
		}
		if c < '0' || c > '9' {
			return 0
		}
		result = result*10 + int(c) - '0'
	}
	return result * sign
}
