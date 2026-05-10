package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func isPalindrome(head *ListNode) bool {
	l1 := head
	cnt := 0

	for l1 != nil {
		cnt++
		l1 = l1.Next

	}
	if cnt == 1 {
		return true
	}
	s1 := make([]int, cnt)
	l1 = head
	if cnt%2 == 0 {
		for i := 0; i < cnt/2; i++ {
			s1[i] = l1.Val
			l1 = l1.Next
		}
		k := cnt/2 - 1
		for l1 != nil {

			if l1.Val != s1[k] {
				return false
			}
			k--
			l1 = l1.Next
		}
	}
	if cnt%2 == 1 {
		for i := 0; i < cnt/2; i++ {
			s1[i] = l1.Val
			l1 = l1.Next
		}
		l1 = l1.Next
		k := cnt/2 - 1
		for l1 != nil {

			if l1.Val != s1[k] {
				return false
			}
			k--
			l1 = l1.Next
		}
	}
	return true
}
