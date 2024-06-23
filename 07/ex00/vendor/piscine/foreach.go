package piscine

import "ft"

func PrintNbr(n int) {
	if n < 0 {
		ft.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		PrintNbr(n / 10)
	}
	ft.PrintRune(rune(n%10) + '0')
}

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}
