package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	// Count the number of elements in the list
	count := 0
	it := l
	for it != nil {
		count++
		it = it.Next
	}
	// Create a slice to store the elements
	slice := make([]int, count)
	it = l
	for i := 0; i < count; i++ {
		slice[i] = it.Data
		it = it.Next
	}
	// Sort the slice
	for i := 0; i < count; i++ {
		for j := i; j < count; j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	// Create a new list with the sorted elements
	newList := &NodeI{Data: slice[0]}
	it = newList
	for i := 1; i < count; i++ {
		it.Next = &NodeI{Data: slice[i]}
		it = it.Next
	}
	return newList
}
