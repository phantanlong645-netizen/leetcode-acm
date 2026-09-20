package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  0,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  -4,
		Next: nil,
	}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l2
	res := detectCycle(l1)
	fmt.Println(res.Val)

}
func detectCycle(head *ListNode) *ListNode {
	l1 := head
	m := map[*ListNode]int{}
	for l1 != nil {
		if _, ok := m[l1]; ok {
			return l1
		}

		m[l1] = 1
		l1 = l1.Next
	}
	return nil

}
