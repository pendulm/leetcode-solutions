package main

import (
    "fmt"
)


func missingNumber(nums []int) int {
    sum := 0
    length := 0
    for _, n := range(nums) {
        sum += n
        length += 1
    }
    return (length+1) * length / 2 - sum
}


func main() {
    nums := []int{3,0,1}
    fmt.Println(missingNumber(nums))
    nums = []int{9,6,4,2,3,5,7,0,1}
    fmt.Println(missingNumber(nums))
}
