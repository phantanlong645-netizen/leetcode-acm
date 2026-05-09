package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	L := map[*ListNode]bool{}
	LA := headA
	for LA != nil {
		L[LA] = true
		LA = LA.Next
	}
	LB := headB
	for LB != nil {
		if L[LB] {
			return LB
		}
		LB = LB.Next
	}
	return nil
}
