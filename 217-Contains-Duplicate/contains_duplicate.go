package main

import (
    "fmt"
)

func containsDuplicate(nums []int) bool {
    seenNums := make(map[int]int)

    for _, k := range(nums) {
        seenNums[k] += 1
    }

    for _, v := range(seenNums) {
        if v > 1 {
            return true
        }
    }
    return false

}

func main() {
    nums := []int{1,2,3,1}
    fmt.Println(containsDuplicate(nums))
    nums = []int{1,2,3,4}
    fmt.Println(containsDuplicate(nums))
    nums = []int{1,1,1,3,3,4,3,2,4,2}
    fmt.Println(containsDuplicate(nums))
}
