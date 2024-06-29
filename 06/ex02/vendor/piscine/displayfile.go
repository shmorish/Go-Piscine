package piscine

import (
	"os"
)

func PrintStr(s string) {
	for _, r := range s {
		os.Stdout.WriteString(string(r))
	}
}

func DisplayFile(s string) {
	file, err := os.Open(s)
	if err != nil {
		os.Stderr.WriteString("open " + s + ": no such file or directory\n")
		return
	}
	defer file.Close()
	buf := []byte{0}
	for {
		_, err := file.Read(buf)
		if err != nil {
			break
		}
		os.Stdout.Write(buf)
	}
}
