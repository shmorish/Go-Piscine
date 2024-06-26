package piscine

import (
	"os"
)

func PrintStr(s string) {
	for _, r := range s {
		os.Stdout.WriteString(string(r))
	}
}

func PrintErr(s string) {
	for _, r := range s {
		os.Stderr.WriteString(string(r))
	}
}

func VecLen(s []string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

const FileNotFound = "File name missing\n"
const TooManyArguments = "Too many arguments\n"

func Cat(args []string) int {
	if VecLen(args) == 1 {
		PrintErr(FileNotFound)
		return 1
	} else if VecLen(args) >= 3 {
		PrintErr(TooManyArguments)
		return 1
	}
	file, err := os.Open(args[1])
	if err != nil {
		PrintErr("open " + args[1] + ": no such file or directory\n")
		return 1
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
	return 0
}
