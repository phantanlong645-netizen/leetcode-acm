package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a1 := &ListNode{
		Val:  2,
		Next: nil,
	}
	a2 := &ListNode{
		Val:  4,
		Next: nil,
	}
	a3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	a1.Next = a2
	a2.Next = a3
	b1 := &ListNode{
		Val:  5,
		Next: nil,
	}
	b2 := &ListNode{
		Val:  6,
		Next: nil,
	}
	b3 := &ListNode{
		Val:  4,
		Next: nil,
	}
	b1.Next = b2
	b2.Next = b3
	c := addTwoNumbers(a1, b1)
	for c != nil {
		fmt.Println(c.Val)
		c = c.Next
	}

}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var n1 int
	var n int
	l3 := &ListNode{}
	head := l3
	n = 0
	for l1 != nil || l2 != nil {
		sum := l1.Val + l2.Val + n
		n1 = sum % 10
		n = sum / 10
		tmp := &ListNode{
			Val:  n1,
			Next: nil,
		}
		l3.Next = tmp
		l3 = l3.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil {
		k1 := l1
		for k1 != nil {
			sum := k1.Val + n
			n = sum / 10
			n2 := sum % 10
			tmp := &ListNode{
				Val:  n2,
				Next: nil,
			}
			l3.Next = tmp
			l3 = l3.Next
			k1 = k1.Next
		}
	}
	if l2 != nil {
		k1 := l1
		for k1 != nil {
			sum := k1.Val + n
			n = sum / 10
			n2 := sum % 10
			tmp := &ListNode{
				Val:  n2,
				Next: nil,
			}
			l3.Next = tmp
			l3 = l3.Next
			k1 = k1.Next
		}
	}
	return head.Next
}
