package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Print(findAnagrams(s, p))
}
func findAnagrams(s string, p string) []int {
	length := len(p)
	str := [26]int{}
	ptr := [26]int{}
	left := 0
	ans := []int{}
	for i := 0; i < length; i++ {
		ptr[p[i]-'a']++
	}
	for i, _ := range s {
		k := s[i] - 'a'
		str[k]++
		left = i - length + 1
		if left < 0 {
			continue
		}
		if str == ptr {
			ans = append(ans, left)
		}

		str[s[left]-'a']--
	}
	return ans
}
