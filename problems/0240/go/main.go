package main

func main() {

}
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix) - 1
	n := len(matrix[0]) - 1

	row := 0
	for true {
		if row > m || n < 0 {
			return false
		}
		if matrix[row][n] == target {
			return true
		}

		if matrix[row][n] < target {
			row++
		} else if matrix[row][n] > target {
			n--
		}

	}
	return false
}
