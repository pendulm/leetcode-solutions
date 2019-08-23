package main

import (
	"fmt"
)

func sumMatrix(matrix [][]int, x int, y int, w int, h int) int {
	sum := 0
	for j := x; j < x+w; j++ {
		for i := y; i < y+h; i++ {
			sum += matrix[i][j]
		}
	}
	return sum
}

func numSubmatrixSumTarget2(matrix [][]int, target int) int {
	row := len(matrix)
	col := len(matrix[0])
	count := 0

	for w := 1; w <= col; w++ {
		for h := 1; h <= row; h++ {
			for x := 0; x < col-w+1; x++ {
				for y := 0; y < row-h+1; y++ {
					sum := sumMatrix(matrix, x, y, w, h)
					if sum == target {
						count += 1
					}
				}
			}

		}
	}
	return count

}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	row := len(matrix)
	col := len(matrix[0])
	count := 0

	for i := 0; i < row; i++ {
		prefixSum := 0
		for j := 0; j < col; j++ {
			matrix[i][j] += prefixSum
			prefixSum = matrix[i][j]
		}
	}

	for i := 0; i < col; i++ {
		for j := i; j < col; j++ {
			seqSum := 0
			targetCount := make(map[int]int)
			targetCount[seqSum] += 1

			for k := 0; k < row; k++ {
				if i == j {
					seqSum += matrix[k][j]
				} else {
					seqSum += matrix[k][j] - matrix[k][i]
				}
				count += targetCount[seqSum-target]
				targetCount[seqSum] += 1
			}

		}
	}
	return count

}

func main() {
	matrix := [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}
	target := 0
	fmt.Println(numSubmatrixSumTarget(matrix, target))
	matrix = [][]int{{1, -1}, {-1, 1}}
	target = 0
	fmt.Println(numSubmatrixSumTarget(matrix, target))
	matrix = [][]int{{0, 1, 1, 1, 0, 1}, {0, 0, 0, 0, 0, 1}, {0, 0, 1, 0, 0, 1}, {1, 1, 0, 1, 1, 0}, {1, 0, 0, 1, 0, 0}}
	target = 0
	fmt.Println(numSubmatrixSumTarget(matrix, target))
	matrix = [][]int{{0, 1, 0, 1, 0, 0}, {0, 0, 0, 0, 1, 0}, {0, 0, 1, 1, 1, 1}, {1, 0, 1, 1, 1, 0}, {0, 1, 1, 0, 0, 0}}
	target = 1
	fmt.Println(numSubmatrixSumTarget(matrix, target))
}
