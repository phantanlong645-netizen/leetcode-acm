package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	t1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	t2 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	t3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	t1.Right = t2
	t2.Left = t3
	ans := make([]int, 0)
	inorder(t1, &ans)
	fmt.Println(ans)

}
func inorder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, ans)
	*ans = append(*ans, root.Val)
	inorder(root.Right, ans)
}
