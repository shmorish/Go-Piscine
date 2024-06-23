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
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		os.Stdout.Write(buf[:n])
		if err != nil {
			break
		}
	}
}
