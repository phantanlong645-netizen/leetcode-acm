package main

import "fmt"

type ListNode struct {
	val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		val:  3,
		Next: nil,
	}
	l2 := &ListNode{
		val:  2,
		Next: nil,
	}
	l3 := &ListNode{
		val:  0,
		Next: nil,
	}
	l4 := &ListNode{
		val:  -4,
		Next: nil,
	}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l2
	res := hasCycle(l1)
	fmt.Println(res)
}

func hasCycle(head *ListNode) bool {
	m := map[*ListNode]bool{}
	l1 := head
	for l1 != nil {
		if _, ok := m[l1]; ok {
			return true
		}
		m[l1] = true
		l1 = l1.Next
	}
	return false

}
