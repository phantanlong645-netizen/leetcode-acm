package main

import (
	"fmt"
)

type ListNode struct {
	value int
	Next  *ListNode
}

func main() {
	l1 := &ListNode{
		value: 1,
		Next:  nil,
	}
	l2 := &ListNode{
		value: 2,
		Next:  nil,
	}
	l3 := &ListNode{
		value: 3,
		Next:  nil,
	}
	l4 := &ListNode{
		value: 1,
		Next:  nil,
	}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	res := isPalindrome(l1)
	fmt.Println(res)

}
func isPalindrome(head *ListNode) bool {
	var cnt int
	cnt = 0
	l1 := head
	for l1 != nil {
		cnt++
		l1 = l1.Next
	}
	l1 = head
	var m []*ListNode
	if cnt%2 == 0 {
		for i := 1; i <= cnt/2; i++ {
			m = append(m, l1)
			l1 = l1.Next
		}
		j := 0
		for l1 != nil {
			if l1.value != m[cnt/2-j-1].value {
				return false
			}
			l1 = l1.Next
			j++
		}
	} else {
		for i := 1; i <= cnt/2; i++ {
			m = append(m, l1)
			l1 = l1.Next
		}
		l1 = l1.Next
		j := 0
		for l1 != nil {
			if l1.value != m[cnt/2-j-1].value {
				return false
			}
			l1 = l1.Next
			j++
		}
	}
	return true

}
