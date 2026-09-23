package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return search(root.Left, root.Right)
}
func search(l, r *TreeNode) bool {
	if l == nil || r == nil {
		return l == r
	}
	return l.Val == r.Val && search(l.Left, r.Right) && search(l.Right, r.Left)
}
