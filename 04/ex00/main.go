package main

import (
	"ft"
	"piscine"
)

func printStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func main() {
	printStr(piscine.PrintProgramname())
}
