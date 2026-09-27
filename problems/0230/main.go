package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	ans := 0
	return inorder(root, &k, &ans)
}
func inorder(root *TreeNode, k *int, ans *int) int {
	if root == nil {
		return 0
	}
	inorder(root.Left, k, ans)
	*k--
	if *k == 0 {
		*ans = root.Val
	}
	inorder(root.Right, k, ans)
	return *ans

}
