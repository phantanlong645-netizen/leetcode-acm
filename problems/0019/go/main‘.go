package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l1 := head
	ans := 0
	for l1 != nil {
		ans++
		l1 = l1.Next
	}
	if ans == 1 {
		return nil
	}
	l1 = head
	k := ans - n
	if ans == n {
		return head.Next
	}
	for i := 0; i < k-1; i++ {
		l1 = l1.Next
	}
	l2 := l1.Next
	l1.Next = l2.Next
	return head
}
