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

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	for i := 0; i < pos; i++ {
		if l == nil {
			return nil
		}
		l = l.Next
	}
	return l
}
