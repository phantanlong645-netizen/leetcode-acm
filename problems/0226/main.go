package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {

}
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	tmp1 := root.Left
	tmp2 := root.Right
	root.Right = tmp1
	root.Left = tmp2
	return root
}
