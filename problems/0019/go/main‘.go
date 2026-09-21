package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	n := 2
	l := removeNthFromEnd(l1, n)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := head
	var cnt int
	cnt = 0
	for l != nil {
		cnt++
		l = l.Next
	}
	l = head
	if cnt == n {
		return head.Next
	}
	for i := 0; i < cnt-n-1; i++ {
		l = l.Next
	}
	tmp := l.Next
	l.Next = tmp.Next
	return head

}
