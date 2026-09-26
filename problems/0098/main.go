package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	t1 := &TreeNode{
		Val: 2,
	}
	t2 := &TreeNode{
		Val: 1,
	}
	t3 := &TreeNode{
		Val: 3,
	}
	t1.Left = t2
	t1.Right = t3
	res := isValidBST(t1)
	fmt.Println(res)
}
func isValidBST(root *TreeNode) bool {
	nums := make([]int, 0)
	inorder(root, &nums)
	pre := math.MinInt64
	flag := true
	for _, v := range nums {
		if v <= pre {
			flag = false
			break
		}
		pre = v
	}
	return flag

}
func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right, nums)

}
