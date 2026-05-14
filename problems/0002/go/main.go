package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	res := pre
	n := 0
	sum := 0
	n1 := 0
	n2 := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {
			n1 = 0
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		} else {
			n2 = 0
		}
		sum = (n1 + n2 + n) % 10
		n = (n1 + n2 + n) / 10
		cur := &ListNode{Val: sum}
		pre.Next = cur
		pre = cur
	}
	if n > 0 {
		cur := &ListNode{Val: n}
		pre.Next = cur

	}
	return res.Next
}
