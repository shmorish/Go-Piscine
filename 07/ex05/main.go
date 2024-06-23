package main

import (
	"fmt"
	"os"
	"piscine"
)

// Write a program that is called doop.
// • The program has to be used with three arguments:
// ◦ A value
// ◦ An operator, one of : +, -, /, *, %
// ◦ Another value
// • In case of an invalid operator, value, number of arguments or an overflow, the
// programs prints nothing.
// • The program has to handle the modulo and division operations by 0 as shown on
// the output examples below.

func main() {
	if len(os.Args) != 4 {
		return
	}
	a, errA := piscine.Atoi(piscine.CheckArgs(os.Args[1]))
	b := os.Args[2]
	c, errB := piscine.Atoi(piscine.CheckArgs(os.Args[3]))
	if errA != "" || errB != "" {
		return
	}
	if b == "+" {
		fmt.Println(a + c)
	} else if b == "-" {
		fmt.Println(a - c)
	} else if b == "*" {
		fmt.Println(a * c)
	} else if b == "/" {
		if c == 0 {
			fmt.Println("No division by 0")
		} else {
			fmt.Println(a / c)
		}
	} else if b == "%" {
		if c == 0 {
			fmt.Println("No modulo by 0")
		} else {
			fmt.Println(a % c)
		}
	}
}
