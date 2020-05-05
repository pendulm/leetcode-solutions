package main

import (
	"fmt"
)

func kLengthApart(nums []int, k int) bool {
	if k >= len(nums) {
		return false

	}
	// len >= 1
	if k == 0 {
		return true
	}

	lastPos := -1
	for i, n := range nums {
		if n == 1 {
			// first 1
			if lastPos == -1 {
				lastPos = i
			} else {
				if i-lastPos <= k {
					return false

				} else {
					lastPos = i
				}

			}
		}

	}
	return true

}

func main() {
	nums := []int{1, 0, 0, 0, 1, 0, 0, 1}
	k := 2
	fmt.Println(kLengthApart(nums, k))

	nums = []int{1, 0, 0, 1, 0, 1}
	k = 2
	fmt.Println(kLengthApart(nums, k))

	nums = []int{1, 1, 1, 1, 1}
	k = 0
	fmt.Println(kLengthApart(nums, k))

	nums = []int{0, 1, 0, 1}
	k = 1
	fmt.Println(kLengthApart(nums, k))
}
