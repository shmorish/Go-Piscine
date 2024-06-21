package piscine

func IsPrintable(s string) bool {
	for _, r := range s {
		if r < 32 {
			return false
		}
	}
	return true
}
