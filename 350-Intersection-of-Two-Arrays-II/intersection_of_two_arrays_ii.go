package main

import (
    "fmt"
)

func min(a, b int) int {
    if (a < b) {
        return a
    } else {
        return b
    }
}


func intersection(nums1 []int, nums2 []int) []int {
    set1 := make(map[int]int)
    set2 := make(map[int]int)

    for _, num := range nums1 {
        set1[num] += 1
    }

    for _, num := range nums2 {
        set2[num] += 1
    }

    var arr []int
    for num, c1  := range set1 {
        count := min(c1, set2[num])
        for i := 0; i < count; i++ {
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


    nums1 = []int{1,2,2,2,1}
    nums2 = []int{2,2}
    intersect = intersection(nums1, nums2)
    fmt.Println(intersect)
}
