package main

import "sort"

func main() {

}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i, _ := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		L := i + 1
		R := len(nums) - 1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				for nums[L] == nums[L+1] {
					L++
				}
				for nums[R-1] == nums[R] {
					R--
				}
				ans = append(ans, []int{i, L, R})
				L++
				R--

			} else if sum < 0 {
				L++
			} else {
				R--
			}

		}

	}
	return ans

}
