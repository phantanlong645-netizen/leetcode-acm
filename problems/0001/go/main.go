package main

import "fmt"

func main() {
	a := []int{2, 7, 11, 15}
	target := 9
	fmt.Print(twoSum2(a, target))

}

/*
思路一：暴力枚举

固定一个数 nums[i]，再从它后面枚举 nums[j]，
如果 nums[i] + nums[j] == target，说明找到了答案，返回下标 [i, j]。

时间复杂度 O(n^2)，空间复杂度 O(1)。
*/
func twoSum1(nums []int, target int) []int {
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

/*
思路二：哈希表

遍历数组时，假设当前数字是 v，需要找的另一个数就是 target - v。
用哈希表记录已经遍历过的数字和它们的下标。

每次遍历到 v 时，先去哈希表里找 target-v：
如果存在，说明之前的某个数和当前 v 相加正好等于 target，直接返回两个下标；
如果不存在，就把当前数字 v 和下标 i 存入哈希表，留给后面的数字匹配。

这样只需要遍历一遍数组，时间复杂度 O(n)，空间复杂度 O(n)。
*/
func twoSum2(nums []int, target int) []int {
	Map := map[int]int{}
	for i, v := range nums {
		if j, ok := Map[target-v]; ok {
			return []int{j, i}
		}
		Map[v] = i
	}
	return nil
}
