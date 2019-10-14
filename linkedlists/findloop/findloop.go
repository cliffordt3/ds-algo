package findloop

type LinkedList struct {
	value int
	next  *LinkedList
}

func FindLoop(head *LinkedList) *LinkedList {
	first := head.next
	second := head.next.next 
	for first != second {
		first = first.next
		second = second.next.next
	}
	first = head
	for first != second {
		first = first.next
		second = second.next
	}
	return first
}
