package main

func main() {

}
func firstMissingPositive(nums []int) int {
	le := len(nums)
	for i, _ := range nums {
		for nums[i] > 0 && nums[i] <= le && nums[nums[i]-1] != nums[i] {
			temp := nums[i]
			nums[i] = nums[nums[i]-1]
			nums[temp-1] = temp
		}
	}
	for i, v := range nums {
		if i != v-1 {
			return i + 1
		}
	}
	return le + 1
}
