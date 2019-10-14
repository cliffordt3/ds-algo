package mergesortedlists

type LinkedList struct {
	value int
	next  *LinkedList
}

func NewLinkedList(val int) *LinkedList {
	return &LinkedList{
		value: val,
		next:  nil,
	}
}

func mergeSortedLists(l1, l2 *LinkedList) *LinkedList {
	head := NewLinkedList(0)
	current := head
	for l1 != nil && l2 != nil {
		if l1.value < l2.value {
			current.next = l1
			l1 = l1.next
		} else {
			current.next = l2
			l2 = l2.next
		}
		current = current.next
	}
	if l1 != nil {
		current.next = l1
	} else if l2 != nil {
		current.next = l2
	}
	return head.next
}
