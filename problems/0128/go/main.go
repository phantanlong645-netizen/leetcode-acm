package main

import "fmt"

func main() {
	a := []int{100, 4, 200, 1, 3, 2}
	fmt.Print(longestConsecutive(a))

}

/*
思路：哈希表 + 枚举连续序列起点

先用哈希表记录 nums 中所有数字，方便 O(1) 判断某个数是否存在。
然后遍历 nums，只从连续序列的起点开始向后查找。

如果 v-1 存在，说明 v 不是起点，直接跳过；
如果 v-1 不存在，说明 v 是一段连续序列的起点，
就不断判断 v+1、v+2、v+3... 是否存在，直到断开。

假设第一个不存在的数是 y，那么从 v 开始的连续序列长度就是 y - v。
每找到一段序列，就更新答案 ans。

因为只有起点才会向后扩展，非起点都会被跳过，
所以每个数字整体最多被访问常数次，时间复杂度 O(n)。
哈希表需要存储所有数字，空间复杂度 O(n)。
*/
func longestConsecutive(nums []int) int {
	hash := map[int]bool{}
	ans := 0
	for _, v := range nums {
		hash[v] = true
	}
	for _, v := range nums {
		if hash[v-1] {
			continue
		}
		y := v + 1
		for hash[y] {
			y++
		}
		ans = max(ans, y-v)
	}
	return ans
}
