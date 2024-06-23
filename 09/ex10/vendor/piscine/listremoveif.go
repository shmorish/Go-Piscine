package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	it := l.Head
	for it != nil {
		if it.Data == data_ref {
			it = it.Next
			l.Head = it
		} else {
			break
		}
	}
	for it != nil && it.Next != nil {
		if it.Next.Data == data_ref {
			if it.Next.Next != nil {
				it.Next = it.Next.Next
			} else {
				it.Next = nil
				l.Tail = it
			}
		} else {
			it = it.Next
		}
	}
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
