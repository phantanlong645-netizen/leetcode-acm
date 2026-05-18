package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

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
	k := l2
	temp := l2.Next
	l3 := &ListNode{}
	l3.Next = head
	for l1 != nil && l2 != nil {
		l2.Next = l1
		l1.Next = temp
		l3.Next = l2
		l3 = l1
		l1 = temp
		if l1 != nil && l1.Next != nil {
			l2 = l1.Next
			temp = l2.Next
		} else {
			break
		}
	}
	return k

}
