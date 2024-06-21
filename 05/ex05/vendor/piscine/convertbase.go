package piscine

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	ResBase := 0
	for _, value1 := range nbr {
		for index2, value2 := range baseFrom {
			if value1 == value2 {
				ResBase = ResBase*StrLen(baseFrom) + index2
			}
		}
	}

	x := ""
	for ResBase != 0 {

		x = string(baseTo[ResBase%StrLen(baseTo)]) + x
		ResBase /= StrLen(baseTo)
	}

	return x
}
