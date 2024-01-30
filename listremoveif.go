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
	if l.Head == nil {
		return
	}

	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	curr := l.Head
	for curr != nil && curr.Next != nil {
		if curr.Next.Data == data_ref {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	if l.Head == nil {
		l.Tail = nil
	} else {
		curr = l.Head
		while curr.Next != nil {
			curr = curr.Next
		}
		l.Tail = curr
	}
}