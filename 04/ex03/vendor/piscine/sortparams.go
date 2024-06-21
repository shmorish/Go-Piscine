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

func SortParams(args []string) {
	args = args[1:]
	for i := 0; i < VecLen(args); i++ {
		for j := i + 1; j < VecLen(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
		printStr(args[i])
	}
}
