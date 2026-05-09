package main

func main() {

}
func spiralOrder(matrix [][]int) []int {
	l := 0
	r := len(matrix[0]) - 1
	up := 0
	down := len(matrix) - 1
	ans := []int{}
	for true {
		for i := l; i <= r; i++ {
			ans = append(ans, matrix[up][i])
		}
		up++
		if up > down {
			break
		}

		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][r])
		}
		r--
		if r < l {
			break
		}
		for i := r; i >= l; i-- {
			ans = append(ans, matrix[down][i])
		}
		down--
		if down < up {
			break
		}
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}
	return ans

}
