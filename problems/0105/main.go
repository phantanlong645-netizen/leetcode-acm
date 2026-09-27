package main

import "slices"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {

}
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	index := slices.Index(inorder, preorder[0])
	left := buildTree(preorder[1:1+index], inorder[:index])
	right := buildTree(preorder[1+index:], inorder[index+1:])
	return &TreeNode{
		Val:   preorder[0],
		Left:  left,
		Right: right,
	}
}
