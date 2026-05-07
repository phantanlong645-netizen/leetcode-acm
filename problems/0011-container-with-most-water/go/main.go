package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	height := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &height[i])
	}

	fmt.Fprintln(out, maxArea(height))
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0

	for left < right {
		area := (right - left) * minInt(height[left], height[right])
		res = maxInt(res, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
