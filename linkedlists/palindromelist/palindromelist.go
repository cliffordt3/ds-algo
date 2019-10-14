package palindromelist

type LinkedList struct {
	value int
	next  *LinkedList
}

func TestPalindrome(head *LinkedList) bool {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	first := head
	second := reverse(slow)
	for first != nil && second != nil {
		if first.value != second.value {
			return false
		}
		first = first.next
		second = second.next
	}
	return true

}

func reverse(head *LinkedList) *LinkedList {
	var p1, p2 *LinkedList = nil, head

	for p2 != nil {
		p3 := p2.next
		p2.next = p1
		p1 = p2
		p2 = p3
	}
	return p1
}
