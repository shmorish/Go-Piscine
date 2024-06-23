package piscine

func Comcheck(s string, match []string) bool {
	for _, arg := range match {
		if s == arg {
			return true
		}
	}
	return false
}
