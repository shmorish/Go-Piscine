package piscine

import (
	"os"
)

func PrintStr(s string) {
	for _, r := range s {
		os.Stdout.WriteString(string(r))
	}
}

func VecLen(s []string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func Cat(args []string) int {
	if VecLen(args) == 1 {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			os.Stdout.Write(buf[:n])
			if err != nil {
				break
			}
		}
	} else {
		for i := 1; i < VecLen(args); i++ {
			file, err := os.Open(args[i])
			if err != nil {
				os.Stderr.WriteString("open " + args[i] + ": no such file or directory\n")
				os.Exit(1)
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
	}
	return 0
}
