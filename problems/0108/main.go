package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	nums := []int{-2, 1, 4, 6, 9}
	sortedArrayToBST(nums)

}
func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	m := l / 2
	return &TreeNode{
		Val:   nums[m],
		Left:  sortedArrayToBST(nums[:m]),
		Right: sortedArrayToBST(nums[m+1:]),
	}
}
