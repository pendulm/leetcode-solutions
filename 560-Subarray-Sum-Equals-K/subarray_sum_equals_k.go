package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	targetCount := make(map[int]int)
	prefixSum := 0
	targetCount[prefixSum] += 1
	occur := 0

	for i := 0; i < len(nums); i++ {
		nums[i] += prefixSum
		prefixSum = nums[i]
		occur += targetCount[prefixSum-k]
		targetCount[prefixSum] += 1
	}

	return occur
}

func main() {
	nums := []int{1, 1, 1}
	k := 2
	fmt.Println(subarraySum(nums, k))
}
