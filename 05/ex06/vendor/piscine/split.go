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

func Split(s, sep string) []string {
	var res []string
	for {
		i := Index(s, sep)
		if i == -1 {
			res = append(res, s)
			break
		}
		res = append(res, s[:i])
		sepLen := StrLen(sep)
		s = s[i+sepLen:]
	}
	return res
}
