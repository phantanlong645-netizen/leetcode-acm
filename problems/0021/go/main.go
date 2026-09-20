package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	a2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	a3 := &ListNode{
		Val:  4,
		Next: nil,
	}
	a1.Next = a2
	a2.Next = a3
	b1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	b2 := &ListNode{
		Val:  3,
		Next: nil,
	}
	b3 := &ListNode{
		Val:  4,
		Next: nil,
	}
	b1.Next = b2
	b2.Next = b3
	l := mergeTwoLists(a1, b1)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l1 := list1
	l2 := list2
	var l3 *ListNode
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l3 = l1
		l1 = l1.Next
	} else {
		l3 = l2
		l2 = l2.Next
	}
	tmp := l3
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			tmp = tmp.Next
			l1 = l1.Next
		} else {
			tmp.Next = l2
			tmp = tmp.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return l3
}
