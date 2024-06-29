package main

import (
	"os"
	"piscine"
)

const FileNotFound = "File name missing\n"
const TooManyArguments = "Too many arguments\n"

func main() {
	arg_len := 0
	for range os.Args {
		arg_len++
	}
	if arg_len == 1 {
		piscine.PrintStr(FileNotFound)
		return
	} else if arg_len >= 3 {
		piscine.PrintStr(TooManyArguments)
		return
	}
	piscine.DisplayFile(os.Args[1])
}
