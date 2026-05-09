package main

func main() {

}
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(0, len(nums)-1, nums)
	reverse(0, k-1, nums)
	reverse(k, len(nums)-1, nums)

}
func reverse(i int, j int, nums []int) {
	temp := 0
	for i < j {
		temp = nums[i]
		nums[i] = nums[j]
		i++
		nums[j] = temp
		j--
	}

}
