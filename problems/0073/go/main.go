package main

func main() {

}
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	rowZero := make([]bool, m)
	colZero := make([]bool, n)
	for i, v := range matrix {
		for j, _ := range v {
			if matrix[i][j] == 0 {
				rowZero[i] = true
				colZero[j] = true
			}
		}
	}
	for i, v := range rowZero {
		for j, v2 := range colZero {
			if v || v2 {
				matrix[i][j] = 0
			}
		}
	}

}
