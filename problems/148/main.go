package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func sortList(head *ListNode) *ListNode {
	cnt := 0
	l := head
	for head != nil {
		l = l.Next
		cnt++
	}
	if cnt == 0 {
		return nil
	}
	if cnt == 1 {
		return head
	}
	l = head
	half := cnt / 2
	for j := 1; j < half; j++ {
		l = l.Next
	}
	l2 := l.Next
	l.Next = nil
	k1 := sortList(head)
	k2 := sortList(l2)
	return mergeList(k1, k2)
}
func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			dummy.Next = l1
			l1 = l1.Next
		} else {
			dummy.Next = l2
			l2 = l2.Next
		}
		dummy = dummy.Next
	}
	if l1 != nil {
		dummy.Next = l1
	}
	if l2 != nil {
		dummy.Next = l2
	}
	return cur.Next
}
