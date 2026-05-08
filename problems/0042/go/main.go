package main

func main() {

}
func trap(height []int) int {
	l := len(height)
	leftmax := make([]int, len(height))
	rightmax := make([]int, len(height))
	leftmax[0] = height[0]
	ans := 0
	rightmax[l-1] = height[len(height)-1]

	for i := 1; i < len(height); i++ {
		leftmax[i] = max(leftmax[i-1], height[i])
	}
	for j := l - 2; j >= 0; j-- {
		rightmax[j] = max(rightmax[j+1], height[j])
	}
	for i, v := range height {
		mini := min(leftmax[i], rightmax[i])
		if mini > v {
			ans += mini - v
		}

	}
	return ans

}
