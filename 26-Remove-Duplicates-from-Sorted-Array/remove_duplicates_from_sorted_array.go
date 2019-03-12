package main

import (
    "fmt"
)

func removeDuplicates(nums []int) int {
    uniqLen := 0
    for i, c := range(nums) {
        if (i == 0) {
            uniqLen = 1
        } else {
            if c != nums[i-1] {
                nums[uniqLen] = c
                uniqLen += 1
            }
        }
    }
    return uniqLen
}

func PrintArrWithLength(nums []int, n int) {
    fmt.Println(nums[0:n])
}

func main() {
    arr := []int{1,1,2}
    n := removeDuplicates(arr)
    PrintArrWithLength(arr, n)
    arr = []int{0,0,1,1,1,2,2,3,3,4}
    n = removeDuplicates(arr)
    PrintArrWithLength(arr, n)
}
