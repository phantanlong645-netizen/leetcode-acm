package main

func main() {

}
func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums)+1)
	suf := make([]int, len(nums)+1)
	pre[0] = nums[0]
	suf[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] * nums[i]
	}
	for j := len(nums) - 2; j >= 0; j-- {
		suf[j] = suf[j+1] * nums[j]
	}
	ans := make([]int, len(nums))
	ans[0] = suf[1]
	l := len(nums)
	ans[l-1] = pre[l-2]
	for i := 1; i < len(nums)-1; i++ {
		ans[i] = pre[i-1] * suf[i+1]
	}
	return ans

}
