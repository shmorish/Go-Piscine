package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
}

func ListReverse(l *List) {
	var prev *NodeL
	current := l.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head, l.Tail = prev, l.Head
}
