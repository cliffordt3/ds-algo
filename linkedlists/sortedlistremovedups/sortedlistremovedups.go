package sortedlistremovedups

type LinkedList struct {
	value int
	next  *LinkedList
}

func RemoveDups(head *LinkedList) *LinkedList {
	current := head
	for current != nil {
		runner := current.next
		for runner != nil && runner.value == current.value {
			runner = runner.next
		}
		current.next = runner
		current = runner
	}
	return head
}
