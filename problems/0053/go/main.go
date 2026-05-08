package main

import "fmt"

func main() {
	s := []int{1, 2}
	fmt.Print(maxSubArray(s))

}
func maxSubArray(nums []int) int {
	cnt := make([]int, len(nums))
	cnt[0] = nums[0]
	ans := cnt[0]
	for i := 0; i < len(nums)-1; i++ {
		if cnt[i]+nums[i+1] > nums[i+1] {
			cnt[i+1] = cnt[i] + nums[i+1]
		} else {
			cnt[i+1] = nums[i+1]
		}
		ans = max(ans, cnt[i+1])
	}
	return ans
}
