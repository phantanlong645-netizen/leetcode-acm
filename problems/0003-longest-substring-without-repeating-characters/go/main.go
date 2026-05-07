package main

func main() {
	s := "abcabcbb"
	println(lengthOfLongestSubstring(s))

}
func lengthOfLongestSubstring(s string) int {
	cnt := [128]int{}
	left := 0
	res := 0
	for right, v := range s {
		cnt[v]++
		for cnt[v] > 1 {
			cnt[s[left]]--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}
