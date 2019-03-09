package main

import (
    "fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
    set1 := make(map[int]int)
    set2 := make(map[int]int)

    for _, num := range nums1 {
        set1[num] = 1
    }

    for _, num := range nums2 {
        set2[num] = 1
    }

    var arr []int
    for num, _ := range set1 {
        _, err := set2[num]
        if err != false {
            arr = append(arr, num)
        }
    }
    return arr
}

func main() {
    nums1 := []int{4,9,5}
    nums2 := []int{9,4,9,8,4}

    intersect := intersection(nums1, nums2)
    fmt.Println(intersect)
}
