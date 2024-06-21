package piscine

func TableLength(table []int) int {
	length := 0
	for range table {
		length++
	}
	return length
}

func SortIntegerTable(table []int) {
	for i := 0; i < TableLength(table); i++ {
		for j := i + 1; j < TableLength(table); j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
