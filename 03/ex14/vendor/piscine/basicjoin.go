package piscine

func BasicJoin(strs []string) string {
	var res string
	for _, s := range strs {
		res += s
	}
	return res
}
