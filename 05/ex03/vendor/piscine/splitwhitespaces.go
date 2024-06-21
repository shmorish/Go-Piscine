package piscine

func SplitWhiteSpaces(str string) []string {
	var result []string
	var word string
	for _, c := range str {
		if c == ' ' || c == '\t' || c == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(c)
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
