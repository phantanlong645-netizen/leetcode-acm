package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l1 := list1
	l2 := list2
	l3 := &ListNode{}
	head := l3
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			l3.Next = l1
			l3 = l1
			l1 = l1.Next
		} else {
			l3.Next = l2
			l3 = l2
			l2 = l2.Next
		}

	}
	if l1 == nil {
		l3.Next = l2
	}
	if l2 == nil {
		l3.Next = l1
	}
	return head.Next

}
