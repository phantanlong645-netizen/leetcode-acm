package main

type LRUCache struct {
	dummy    *ListNode
	capacity int
	m        map[int]*ListNode
}
type ListNode struct {
	key, value int
	pre, next  *ListNode
}

func Constructor(capacity int) LRUCache {
	dummy1 := &ListNode{}
	dummy1.pre = dummy1
	dummy1.next = dummy1
	LRUCache1 := LRUCache{
		dummy:    dummy1,
		capacity: capacity,
		m:        make(map[int]*ListNode, capacity),
	}
	return LRUCache1

}
func (this *LRUCache) remove(num int) {
	node := this.m[num]
	node.pre.next = node.next
	node.next.pre = node.pre
}
func (this *LRUCache) pushFront(node *ListNode) {
	node.pre = this.dummy
	node.next = this.dummy.next
	this.dummy.next.pre = node
	this.dummy.next = node

}
func (this *LRUCache) getNode(num int) *ListNode {
	node := this.m[num]
	if node == nil {
		return nil
	}
	this.remove(num)
	this.pushFront(node)
	return node
}

func (this *LRUCache) Get(key int) int {
	node := this.getNode(key)
	if node != nil {
		return node.value
	} else {
		return -1
	}

}

func (this *LRUCache) Put(key int, value int) {
	node := this.getNode(key)
	if node != nil {
		node.value = value
		return
	} else {
		newNode := &ListNode{
			key:   key,
			value: value,
		}
		this.m[key] = newNode
		this.pushFront(newNode)

		if len(this.m) > this.capacity {
			k := this.dummy.pre
			this.remove(k.key)
			delete(this.m, k.key)

		}
	}
}
