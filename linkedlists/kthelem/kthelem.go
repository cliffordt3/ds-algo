package kthelem

type LinkedList struct {
	value int
	next  *LinkedList
}

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	counter,fast, slow := 1, head, head
	for counter <= k {
		fast = fast.next
		counter++
	}
	if fast == nil {
		head.value = head.next.value
		head.next = head.next.next
		return 
	}
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
	
}
