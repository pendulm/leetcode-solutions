package main

import (
    "fmt"
)

func reverse(arr []int) {
    length := len(arr)
    if (length > 1) {
        for i := 0; i < length / 2; i++ {
            arr[i], arr[length-1-i] = arr[length-1-i], arr[i]
        }
    }
}

func rotate(nums []int, k int)  {
    length := len(nums)
    k = k % length
    reverse(nums)
    reverse(nums[0:k])
    reverse(nums[k:])
}


func main() {
    nums := []int{1,2,3,4,5,6,7}
    rotate(nums, 13)
    fmt.Println(nums)
}
