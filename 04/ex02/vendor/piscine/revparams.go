package piscine

import "ft"

func printStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func VecLen(s []string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func RevParams(args []string) {
	for i := VecLen(args) - 1; i > 0; i-- {
		printStr(args[i])
	}
}
