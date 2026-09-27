package main

import "fmt"

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
	t1.Left = t2
	t1.Right = t3
	t2.Right = t5
	t3.Right = t4
	s := rightSideView(t1)
	fmt.Println(s)
}
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	slice := make([]*TreeNode, 0)
	slice = append(slice, root)
	for len(slice) != 0 {
		l := len(slice)
		k := l
		for i := 1; i <= l; i++ {
			if k == 1 {
				ans = append(ans, slice[0].Val)
			}
			node := slice[0]
			slice = slice[1:]
			k--
			if node.Left != nil {
				slice = append(slice, node.Left)
			}
			if node.Right != nil {
				slice = append(slice, node.Right)
			}
		}

	}
	return ans
}
