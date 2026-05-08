package main

func main() {

}

/*
思路：滑动窗口

先统计 t 中每个字符需要出现的次数 need。
然后用左右指针维护 s 中的一个窗口 [left, right)。

right 负责扩大窗口，把字符加入 window；
left 负责在窗口已经包含 t 的所有字符后，尽量缩小窗口。

valid 表示有多少种字符已经满足 need 的要求。
当 valid == len(need) 时，说明当前窗口已经覆盖了 t，
此时更新最短答案，并尝试右移 left 缩小窗口。

时间复杂度 O(n)，空间复杂度 O(字符集大小)。
*/
func minWindow(s string, t string) string {
	t1 := map[byte]int{}
	s1 := map[byte]int{}
	valid := 0
	for i, _ := range t {
		t1[t[i]]++
	}
	ans := 99999999
	left := 0
	right := 0
	str := ""
	for right < len(s) {
		s1[s[right]]++
		if t1[s[right]] != 0 && s1[s[right]] == t1[s[right]] {
			valid++
		}
		for valid == len(t1) {
			if right-left+1 < ans {
				ans = min(ans, right-left+1)
				str = s[left : right+1]
			}
			c := s[left]
			if t1[c] != 0 && t1[c] == s1[c] {
				valid--
			}
			s1[c]--
			left++
		}
		right++

	}
	return str
}
