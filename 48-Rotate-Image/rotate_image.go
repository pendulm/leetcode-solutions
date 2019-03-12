package main

import (
    "fmt"
)

func reverseBydiagonal(matrix [][]int) {
    k := len(matrix)
    for i := 0; i < k; i++ {
        for j := i+1; j < k; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
}


func reverseByColumn(matrix [][]int) {
    k := len(matrix)
    for j := 0; j < k/2; j++ {
        for i := 0; i < k; i++ {
            matrix[i][j], matrix[i][k-1-j] = matrix[i][k-1-j], matrix[i][j]
        }
    }
}

func rotate(matrix [][]int)  {
    reverseBydiagonal(matrix)
    reverseByColumn(matrix)
}


func main() {
    matrix := [][]int{ {1,2,3}, {4,5,6}, {7,8,9} }
    rotate(matrix)
    fmt.Println(matrix)
    matrix = [][]int{ { 5, 1, 9,11}, { 2, 4, 8,10}, {13, 3, 6, 7}, {15,14,12,16}}
    rotate(matrix)
    fmt.Println(matrix)
}
