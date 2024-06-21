package piscine

func VecLen(s []string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func ConcatParams(args []string) string {
	var str string
	line := "\n"
	for i := 0; i < VecLen(args); i++ {
		if i != VecLen(args)-1 {
			str += (args[i] + line)
		} else {
			str += args[i]
		}
	}
	return str
}
