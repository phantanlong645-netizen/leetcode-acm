package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	n2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	n3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	n4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	L := swapPairs(n1)
	for L != nil {
		fmt.Println(L.Val)
		L = L.Next
	}

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	l1 := head
	l2 := head.Next
	K := l2
	l3 := &ListNode{}
	l3.Next = l1
	for l1 != nil && l2 != nil {
		tmp := l2.Next
		l2.Next = l1
		l1.Next = tmp
		l3.Next = l2
		l3 = l1
		l1 = tmp
		if l1 != nil && l1.Next != nil {
			l2 = l1.Next
		} else {
			break
		}

	}
	return K
}
