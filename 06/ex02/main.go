package main

import (
	"os"
	"piscine"
)

func main() {
	arg_len := 0
	for range os.Args {
		arg_len++
	}
	if arg_len == 1 {
		piscine.PrintStr("File name missing\n")
		return
	} else if arg_len >= 3 {
		piscine.PrintStr("Too many arguments\n")
		return
	}
	piscine.DisplayFile(os.Args[1])
}
