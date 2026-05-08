package main

func main() {

}
func subarraySum(nums []int, k int) int {
	s := make([]int, len(nums))
	res := 0
	for i, v := range nums {
		s[i+1] = s[i] + v
	}
	cnt := map[int]int{}
	for _, su := range s {
		res += cnt[su-k]
		cnt[su]++
	}
	return res
}
