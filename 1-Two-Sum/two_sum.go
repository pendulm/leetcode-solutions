package main

import (
    "fmt"
)

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if nums[i] + nums[j] == target && i != j {
                return []int{i, j}
            }
        }
    }
    return []int{-1, -1}

}

func twoSumMap(nums []int, target int) []int {
    table := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        j, ok := table[target - nums[i]]
        if ok == true && i != j {
            return []int{j, i}
        }
        table[nums[i]] = i
    }
    return []int{-1, -1}
}

func main() {
    nums := []int{2, 7, 11, 15}
    fmt.Println(twoSum(nums, 9))
    fmt.Println(twoSumMap(nums, 9))
}
