package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {

}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	return r

}
