package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {

}
func levelOrder(root *TreeNode) [][]int {
	q := make([]*TreeNode, 0)
	q = append(q, root)
	res := make([][]int, 0)
	for len(q) != 0 {
		levelSize := len(q)
		ans := make([]int, 0)
		for i := 0; i < levelSize; i++ {
			node := q[0]
			ans = append(ans, node.Val)
			q := q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, ans)

	}
	return res
}
