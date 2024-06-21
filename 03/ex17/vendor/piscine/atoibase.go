package piscine

func StrLen(str string) int {
	var count int
	for range str {
		count++
	}
	return count
}

func AtoiBase(s string, base string) int {
	nbr := 0
	sign := 1

	if StrLen(base) < 2 {
		return 0
	}
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	for _, c := range s {
		nbr *= StrLen(base)
		for i, v := range base {
			if c == v {
				nbr += i
			}
			if v == '-' || v == '+' {
				return 0
			}
		}
	}
	return nbr * sign
}
