package main

import (
	"ft"
	"os"
)

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return true
	} else {
		return false
	}
}

const EvenMsg = "I have an even number of arguments"
const OddMsg = "I have an odd number of arguments"

func main() {
	lengthOfArg := 0
	for range os.Args {
		lengthOfArg++
	}
	if isEven(lengthOfArg) == true {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
