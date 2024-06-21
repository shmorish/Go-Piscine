package piscine

import "ft"

func printStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

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

func PrintWordsTables(a []string) {
	for _, word := range a {
		printStr(word)
	}
}
