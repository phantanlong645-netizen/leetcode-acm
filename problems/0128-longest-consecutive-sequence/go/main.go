package main

import "fmt"

func main() {
	a := []int{100, 4, 200, 1, 3, 2}
	fmt.Print(longestConsecutive(a))

}
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
