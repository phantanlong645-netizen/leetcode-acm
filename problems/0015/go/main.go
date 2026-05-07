package main

import "sort"

func main() {

}

/*
思路：排序 + 双指针

先对数组排序，然后枚举第一个数 nums[i]。
固定 nums[i] 后，剩下的问题就变成：
在 i 右边找两个数 nums[L] 和 nums[R]，使三数之和等于 0。

因为数组已经有序：
如果 sum < 0，说明当前和太小，需要让左指针 L 右移，使和变大；
如果 sum > 0，说明当前和太大，需要让右指针 R 左移，使和变小；
如果 sum == 0，就记录答案，并跳过 L 和 R 位置上的重复数字。

为了避免重复三元组：
1. 枚举 i 时，如果 nums[i] 和 nums[i-1] 相同，直接跳过；
2. 找到一组答案后，移动 L、R 前也要跳过重复值。

时间复杂度 O(n^2)，空间复杂度不算返回结果为 O(1)。
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		L := i + 1
		R := len(nums) - 1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R-1] == nums[R] {
					R--
				}
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
