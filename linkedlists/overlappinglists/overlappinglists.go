package overlappinglists

type LinkedList struct {
	value int
	next  *LinkedList
}

func FindOverlappingNode(l1, l2 *LinkedList) *LinkedList {
	lenList1 := length(l1)
	lenList2 := length(l2)
	if lenList1 > lenList2 {
		advanceListByK(lenList1-lenList2, l1)
	} else if lenList2 > lenList1 {
		advanceListByK(lenList2-lenList1, l2)
	}

	for l1 != nil && l2 != nil && l1 != l2 {
		l1 = l1.next
		l2 = l2.next
	}
	return l1

}

func length(l *LinkedList) int {
	len := 0
	for l != nil {
		l = l.next
		len++
	}
	return len
}

func advanceListByK(k int, l *LinkedList) {
	for k > 0 {
		k = k - 1
		l = l.next
	}
}
