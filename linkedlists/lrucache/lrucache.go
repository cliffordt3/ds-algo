package lrucache

// Do not edit the class below except for the insertKeyValuePair,
// getValueFromKey, and getMostRecentKey methods. Feel free
// to add new properties and methods to the class.
type LRUCache struct {
	index            map[string]*DoublyLinkedListNode
	maxSize          int
	currentSize      int
	listofMostRecent *DoublyLinkedList
	// Add fields here.
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		index:            map[string]*DoublyLinkedListNode{},
		maxSize:          size,
		currentSize:      0,
		listofMostRecent: &DoublyLinkedList{},
	}
}

type DoublyLinkedListNode struct {
	key   string
	value int
	prev  *DoublyLinkedListNode
	next  *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}


func (cache *LRUCache) InsertKeyValuePair(key string, value int) {
	// Write your code here.
	if _, found := cache.index[key]; found {
		cache.replaceKey(key, value)
	} else {
		if cache.currentSize == cache.maxSize {
			cache.evictLeastRecent()
		} else {
			cache.currentSize++
		}
		cache.index[key] = &DoublyLinkedListNode{
			key:   key,
			value: value,
		}

	}
	cache.updateMostRecent(cache.index[key])
}

func (cache *LRUCache) replaceKey(key string, value int) {
	if node, found := cache.index[key]; !found {
		panic("The provided key is not in the cache")
	} else {
		node.value = value
	}
}

func (cache *LRUCache) evictLeastRecent() {
	key := cache.listofMostRecent.tail.key
	cache.listofMostRecent.removeTail()
	delete(cache.index, key)
}

func (cache *LRUCache) updateMostRecent(node *DoublyLinkedListNode) {
	cache.listofMostRecent.setHeadTo(node)
}

// The second return value indicates whether or not the key was found
// in the cache.
func (cache *LRUCache) GetValueFromKey(key string) (int, bool) {
	if node, found := cache.index[key]; !found {
		return 0, false
	} else {
		cache.updateMostRecent(cache.index[key])
		return node.value, true
	}

}

// The second return value is false if the cache is empty.
func (cache *LRUCache) GetMostRecentKey() (string, bool) {
	if cache.listofMostRecent.head == nil {
		return "", false
	}
	return cache.listofMostRecent.head.key, true
}


func (node *DoublyLinkedListNode) removeBindings() {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev, node.next = nil, nil
}


func (list *DoublyLinkedList) setHeadTo(node *DoublyLinkedListNode) {
	if list.head == node {
		return
	}
	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}
	if list.head == list.tail {
		list.tail.prev = node
		list.head = node
		list.head.next = list.tail
		return
	}
	if list.tail == node {
		list.removeTail()
	}
	list.head.prev = node
	node.next = list.head
	list.head = node
}

func (list *DoublyLinkedList) removeTail() {
	if list.tail == nil {
		return
	}
	if list.head == list.tail {
		list.head, list.tail = nil, nil
		return
	}
	list.tail = list.tail.prev
	list.tail.next.removeBindings()
}
