package main

import (
	"os"
	"piscine"
)

func main() {
	match := []string{"42", "piscine", "piscine 42"}
	for _, arg := range os.Args[1:] {
		if piscine.Comcheck(arg, match) {
			piscine.PrintError("Alert!!!")
			break
		}
	}
}
