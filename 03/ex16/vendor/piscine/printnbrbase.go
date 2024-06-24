package piscine

import "ft"

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

func PrintNumBase(num int, base string) {
	mod := StrLen(base)
	i := 0
	if num == 0 {
		ft.PrintRune('0')
		return
	}
	for j := 1; j <= num%mod; j++ {
		i++
	}
	for j := -1; j >= num%mod; j-- {
		i++
	}
	if num/mod != 0 {
		PrintNumBase(num/mod, base)
	}
	ft.PrintRune(rune(base[i]))
	return
}

func PrintNbrBase(n int, base string) {
	if StrLen(base) < 2 || Index(base, "+") >= 0 || Index(base, "-") >= 0 {
		ft.PrintRune('N')
		ft.PrintRune('V')
		return
	}
	for i1, char1 := range base {
		for i2, char2 := range base {
			if i1 != i2 && char1 == char2 {
				ft.PrintRune('N')
				ft.PrintRune('V')
				return
			}
		}
	}
	if n < 0 {
		ft.PrintRune('-')
	}
	PrintNumBase(n, base)
}
