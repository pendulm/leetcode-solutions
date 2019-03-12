package main

import (
    "fmt"
)

func singleNumber(nums []int) int {
    xor := nums[0]
    for i := 1; i < len(nums); i++ {
        xor ^= nums[i]
    }
    return xor
}


func main() {
    arr := []int{2,2,1}
    fmt.Println(singleNumber(arr))
    arr = []int{-4,1,-2,1,-2}
    fmt.Println(singleNumber(arr))
}
