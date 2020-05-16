package main

import (
	"fmt"
)

func longestSubarray(nums []int, limit int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}

	longest := 1

	for i := 0; i < length; i++ {
		max := nums[i]
		min := nums[i]
		continuous := 0
		for j := i; j < length; j++ {
			cur := nums[j]
			if cur > max {
				max = cur
			}
			if cur < min {
				min = cur
			}
			if max-min <= limit {
				continuous++
				if continuous > longest {
					longest = continuous
				}
			} else {
				break
			}

		}

	}
	return longest

}

func main() {
	nums := []int{8, 2, 4, 7}
	limit := 4
	fmt.Println(longestSubarray(nums, limit))

	nums = []int{10, 1, 2, 4, 7, 2}
	limit = 5
	fmt.Println(longestSubarray(nums, limit))

	nums = []int{4, 2, 2, 2, 4, 4, 2, 2}
	limit = 0
	fmt.Println(longestSubarray(nums, limit))
}
