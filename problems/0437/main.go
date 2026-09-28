package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {

}
func pathSum(root *TreeNode, targetSum int) int {
	ans := 0
	dfs1(root, targetSum, &ans)
	return ans
}
func dfs1(root *TreeNode, targetsum int, ans *int) {
	if root == nil {
		return
	}
	dfs2(root, targetsum, ans, root.Val)
	dfs1(root.Left, targetsum, ans)
	dfs1(root.Right, targetsum, ans)

}
func dfs2(root *TreeNode, targetSum int, ans *int, val int) {
	if root == nil {
		return
	}
	if val == targetSum {
		*ans++
	}
	dfs2(root.Left, targetSum, ans, val+root.Left.Val)
	dfs2(root.Right, targetSum, ans, val+root.Right.Val)
}
