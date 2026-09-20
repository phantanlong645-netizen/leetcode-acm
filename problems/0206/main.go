package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func main() {
	k1 := &ListNode{
		value: 1,
		next:  nil,
	}
	k2 := &ListNode{
		value: 2,
		next:  nil,
	}
	k3 := &ListNode{
		value: 3,
		next:  nil,
	}
	k4 := &ListNode{
		value: 4,
		next:  nil,
	}
	k5 := &ListNode{
		value: 5,
		next:  nil,
	}
	k1.next = k2
	k2.next = k3
	k3.next = k4
	k4.next = k5
	head := reverseList(k1)
	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}

}
func reverseList(head *ListNode) *ListNode {
	var dummy *ListNode
	if head == nil {
		return nil
	}
	if head.next == nil {
		return head
	}
	cur := head
	for cur != nil {
		tmp := cur.next
		cur.next = dummy
		dummy = cur
		cur = tmp
	}
	return dummy
}
