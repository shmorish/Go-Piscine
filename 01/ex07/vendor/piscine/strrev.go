package piscine

func StrRev(s string) string {
	var reversed string
	for _, c := range s {
		reversed = string(c) + reversed
	}
	return reversed
}
