package main

import (
    "fmt"
)

const WIDTH int = 9


func checkRow(board [][]byte, row int) bool {
    var numMark [WIDTH+1]int
    var i int
    for col := 0; col < WIDTH; col++ {
        if board[row][col] == '.' {
            i = 0
        } else {
            i = int(board[row][col]) - '0'
        }
        if numMark[i] != 0 && i != 0 {
            return false
        } else {
            numMark[i] = 1
        }
    }
    return true
}

func checkColumn(board [][]byte, col int) bool {
    var numMark [WIDTH+1]int
    var i int
    for row := 0; row < WIDTH; row++ {
        if board[row][col] == '.' {
            i = 0
        } else {
            i = int(board[row][col]) - '0'
        }
        if numMark[i] != 0 && i != 0 {
            return false
        } else {
            numMark[i] = 1
        }
    }
    return true
}

func checkGrid(board [][]byte, grid int) bool {
    var numMark [WIDTH+1]int
    var i int
    x, y := grid % 3 * 3, grid / 3 * 3
    for row := y; row < y+3; row++ {
        for col := x; col < x+3; col++ {
            if board[row][col] == '.' {
                i = 0
            } else {
                i = int(board[row][col]) - '0'
            }
            if numMark[i] != 0 && i != 0 {
                return false
            } else {
                numMark[i] = 1
            }
        }
    }

    return true
}

func isValidSudoku(board [][]byte) bool {
    for i := 0; i < WIDTH; i++ {
        if (checkRow(board, i) == false) {
            return false
        }
    }

    for i := 0; i < WIDTH; i++ {
        if (checkColumn(board, i) == false) {
            return false
        }
    }

    for i := 0; i < WIDTH; i++ {
        if (checkGrid(board, i) == false) {
            return false
        }
    }
    return true
}


func main() {
    board := [][]byte {
        {'5','3','.','.','7','.','.','.','.'},
        {'6','.','.','1','9','5','.','.','.'},
        {'.','9','8','.','.','.','.','6','.'},
        {'8','.','.','.','6','.','.','.','3'},
        {'4','.','.','8','.','3','.','.','1'},
        {'7','.','.','.','2','.','.','.','6'},
        {'.','6','.','.','.','.','2','8','.'},
        {'.','.','.','4','1','9','.','.','5'},
        {'.','.','.','.','8','.','.','7','9'} }

    fmt.Println(isValidSudoku(board))
    board = [][]byte {
        {'8','3','.','.','7','.','.','.','.'},
        {'6','.','.','1','9','5','.','.','.'},
        {'.','9','8','.','.','.','.','6','.'},
        {'8','.','.','.','6','.','.','.','3'},
        {'4','.','.','8','.','3','.','.','1'},
        {'7','.','.','.','2','.','.','.','6'},
        {'.','6','.','.','.','.','2','8','.'},
        {'.','.','.','4','1','9','.','.','5'},
        {'.','.','.','.','8','.','.','7','9'} }
    fmt.Println(isValidSudoku(board))
}
