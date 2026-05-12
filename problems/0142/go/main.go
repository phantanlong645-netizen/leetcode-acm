package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func detectCycle(head *ListNode) *ListNode {
	m := map[*ListNode]int{}
	l := head
	cnt := 0
	for l != nil {
		if _, ok := m[l]; ok {
			return l
		}
		m[l] = cnt
		cnt++
		l = l.Next
	}

	return nil
}
