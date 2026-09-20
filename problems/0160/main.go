package main

import "fmt"

type ListNode struct {
	value int
	Next  *ListNode
}

func main() {
	La1 := &ListNode{
		value: 4,
		Next:  nil,
	}
	La2 := &ListNode{
		value: 1,
		Next:  nil,
	}
	La1.Next = La2
	Lb1 := &ListNode{
		value: 5,
		Next:  nil,
	}
	Lb2 := &ListNode{
		value: 6,
		Next:  nil,
	}
	Lb3 := &ListNode{
		value: 1,
		Next:  nil,
	}
	Lb1.Next = Lb2
	Lb2.Next = Lb3
	Lab1 := &ListNode{
		value: 8,
		Next:  nil,
	}
	Lab2 := &ListNode{
		value: 4,
		Next:  nil,
	}
	Lab3 := &ListNode{
		value: 5,
		Next:  nil,
	}
	Lab1.Next = Lab2
	Lab2.Next = Lab3
	La2.Next = Lab1
	Lb3.Next = Lab1
	temp := getIntersectionNode(La1, Lb1)
	fmt.Print(temp.value)

}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mapA := map[*ListNode]bool{}
	l1 := headA
	l2 := headB
	for l1 != nil {
		mapA[l1] = true
		l1 = l1.Next
	}
	for l2 != nil {
		if mapA[l2] {
			return l2
		}
		l2 = l2.Next
	}
	return nil

}
