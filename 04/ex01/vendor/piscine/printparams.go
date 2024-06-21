package piscine

import "ft"

func printStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func PrintParams(args []string) {
	args = args[1:]
	for _, arg := range args {
		printStr(arg)
	}
}
