package cyclicallyrightshiftlist

type LinkedList struct {
	value int
	next  *LinkedList
}

func CyclicallyRightShiftList(k int, head *LinkedList) *LinkedList {
	n := 0
	tail := head
	for tail != nil {
		tail = tail.next
		n++
	}
	k = k % n
	if k == 0 {
		return head
	}

	stepsToNewHead := n - k
	tail.next = head
	newTail := tail

	for stepsToNewHead > 0 {
		stepsToNewHead--
		newTail = newTail.next
	}
	newHead := newTail.next
	newTail.next = nil
	return newHead

}
