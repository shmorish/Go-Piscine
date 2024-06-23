package piscine

import (
	"os"
)

func CheckArgs(s string) string {
	for _, c := range s {
		if c < '0' || c > '9' {
			os.Exit(0)
		}
	}
	return s
}

func Atoi(s string) (int, string) {
	n := 0
	for _, c := range s {
		if n > (n*10+int(c-'0'))/10 {
			return 0, "error"
		}
		n = n*10 + int(c-'0')
	}
	return n, ""
}
