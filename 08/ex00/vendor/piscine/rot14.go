package piscine

// Hello! How are You?
// Vszzc! Vck ofs Mci?

func Rot14(str string) string {
	runes := []rune(str)
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = r + 14
			if runes[i] > 'z' {
				runes[i] -= 26
			}
		} else if r >= 'A' && r <= 'Z' {
			runes[i] = r + 14
			if runes[i] > 'Z' {
				runes[i] -= 26
			}
		}
	}
	return string(runes)
}
