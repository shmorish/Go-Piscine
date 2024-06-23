package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, v := range *ptr {
		if v != "" {
			(*ptr)[count] = v
			count++
		}
	}
	*ptr = (*ptr)[:count]
	return count
}
