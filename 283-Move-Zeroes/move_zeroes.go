package main

import (
    "fmt"
)

func moveZeroes(nums []int)  {
    nonZeroCount := 0
    for i, n := range(nums) {
        if n != 0 {
            if (i != nonZeroCount) {
                nums[nonZeroCount] = n
            }
            nonZeroCount += 1
        }
    }
    for i := len(nums) -1; i >= nonZeroCount; i-- {
        nums[i] = 0
    }
}

func main() {
    nums := []int{0,1,0,3,12}
    moveZeroes(nums)
    fmt.Println(nums)
    nums = []int{0,0,0,0,0}
    moveZeroes(nums)
    fmt.Println(nums)
}
