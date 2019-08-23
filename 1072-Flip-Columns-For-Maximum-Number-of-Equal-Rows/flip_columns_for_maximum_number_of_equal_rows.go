package main

import (
	"fmt"
)

func arrToStr(arr []int) string {
	str := ""
	for i := 0; i < len(arr); i++ {
		str += string('0' + byte(arr[i]))
	}
	return str
}

func maxEqualRowsAfterFlips(matrix [][]int) int {
	row := len(matrix)
	col := len(matrix[0])
	max := 0
	rowPattern := make(map[string]int)

	for i := 0; i < row; i++ {
		flip := []int{}
		for j := 0; j < col; j++ {
			flip = append(flip, 1-matrix[i][j])
		}
		rowPattern[arrToStr(matrix[i])] += 1
		count := rowPattern[arrToStr(matrix[i])] + rowPattern[arrToStr(flip)]
		if count > max {
			max = count
		}
	}
	return max
}

func main() {
	matrix := [][]int{{0, 1}, {1, 1}}
	fmt.Println(maxEqualRowsAfterFlips(matrix))
	matrix = [][]int{{0, 1}, {1, 0}}
	fmt.Println(maxEqualRowsAfterFlips(matrix))
}
