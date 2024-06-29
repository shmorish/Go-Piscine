package piscine

import (
	"os"
)

func PrintError(s string) {
	os.Stderr.WriteString(s)
	os.Stderr.WriteString("\n")
}

func Comcheck(s string, match []string) bool {
	for _, arg := range match {
		if s == arg {
			return true
		}
	}
	return false
}
