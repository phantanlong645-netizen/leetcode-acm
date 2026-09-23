package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	t1 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	t2 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	t3 := &TreeNode{
		Val:   20,
		Left:  nil,
		Right: nil,
	}
	t4 := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	t5 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	t1.Left = t2
	t1.Right = t3
	t3.Left = t4
	t3.Right = t5
	ans := search(t1, 0)
	fmt.Println(ans)
}
func search(root *TreeNode, k int) int {
	if root == nil {
		return k
	}
	l := search(root.Left, k) + 1
	r := search(root.Right, k) + 1
	if l < r {
		return r
	} else {
		return l
	}
}
