package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	m := map[*ListNode]bool{}
	h := head
	for h != nil {
		if m[h] {
			return true
		}
		m[h] = true
		h = h.Next
	}
	return false
}
