package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	t1 := &TreeNode{
		Val: 1,
	}
	t2 := &TreeNode{
		Val: 2,
	}
	t3 := &TreeNode{
		Val: 3,
	}
	t4 := &TreeNode{
		Val: 4,
	}
	t5 := &TreeNode{
		Val: 5,
	}
	t6 := &TreeNode{
		Val: 6,
	}
	t1.Left = t2
	t1.Right = t5
	t2.Left = t3
	t2.Right = t4
	t5.Right = t6
	flatten(t1)

}
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	tmp := root.Right
	root.Right = root.Left
	root.Left = nil
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = tmp

}
