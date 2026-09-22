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
		Val:  3,
		Next: nil,
	}
	a4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	a5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	res := reverse(a1, 2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}

func reverse(head *ListNode, k int) *ListNode {
	start := head
	end := head
	dummy := &ListNode{}
	dummy.Next = head
	res := dummy
	var i int
	for start != nil {
		for i = 1; i < k; i++ {
			if end != nil {
				end = end.Next
			} else {
				break
			}
		}
		if end == nil {
			break
		}
		next := end.Next
		end.Next = nil
		dummy.Next = reverseOne(start)
		start.Next = next
		dummy = start
		start = next
		end = next
	}
	return res.Next

}
func reverseOne(head *ListNode) *ListNode {
	var dummy *ListNode
	l1 := head
	for l1 != nil {
		l2 := l1.Next
		l1.Next = dummy
		dummy = l1
		l1 = l2
	}
	return dummy
}
