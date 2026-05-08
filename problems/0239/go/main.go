package main

func main() {

}
func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	ans := make([]int, len(nums)-k+1)
	for i, v := range nums {
		for len(q) > 0 && v <= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		left := i - k + 1
		if q[0] == left {
			q = q[1:]
		}
		if left >= 0 {
			ans = append(ans, q[0])
		}
	}
	return ans
}
