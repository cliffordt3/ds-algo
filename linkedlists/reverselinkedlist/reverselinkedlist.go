package reverselinkedlist

type LinkedList struct {
	value int
	next  *LinkedList
}

func ReverseLinkedList(head *LinkedList) *LinkedList {
	// Write your code here.
	var p1, p2 *LinkedList = nil, head
	for p2 != nil {
		p3 := p2.next
		p2.next = p1
		p1 = p2
		p2 = p3
	}
	return p1
}
