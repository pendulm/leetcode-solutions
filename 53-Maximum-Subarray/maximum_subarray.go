package main

import (
	"fmt"
)

func maxSubArray1(nums []int) int {
	length := len(nums)
	max := nums[0]

	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func maxSubArray2(nums []int) int {
	length := len(nums)
	max := nums[0]

	for i := 0; i < length; i++ {
		sum := 0
		for j := i; j < length; j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func maxSubArray3(nums []int) int {
	length := len(nums)
	max := nums[0]
	if length == 1 {
		return max
	}

	mid := length / 2
	firstHalfMax := maxSubArray3(nums[0:mid])
	secondHalfMax := maxSubArray3(nums[mid:])
	if firstHalfMax > secondHalfMax {
		max = firstHalfMax
	} else {
		max = secondHalfMax
	}

	leftMax, leftSum := nums[mid-1], 0
	for i := mid - 1; i >= 0; i-- {
		leftSum += nums[i]
		if leftSum > leftMax {
			leftMax = leftSum
		}
	}
	rightMax, rightSum := nums[mid], 0
	for i := mid; i < length; i++ {
		rightSum += nums[i]
		if rightSum > rightMax {
			rightMax = rightSum
		}
	}
	if max > leftMax+rightMax {
		return max
	} else {
		return leftMax + rightMax
	}
}

func getSign(num int) int {
	if num >= 0 {
		return 1
	} else {
		return -1
	}
}

func maxSubArray4(nums []int) int {
	max := nums[0]
	aggNum := []int{}
	var keepSign, sum int
	keepSign = getSign(nums[0])

	for i := 0; i < len(nums); i++ {
		numSign := getSign(nums[i])
		if keepSign != numSign {
			aggNum = append(aggNum, sum)
			sum = nums[i]
			keepSign = numSign
		} else {
			sum += nums[i]
		}

		if nums[i] > max {
			max = nums[i]
		}

		if i == len(nums)-1 {
			aggNum = append(aggNum, sum)
		}
	}

	skipLeadingNegetive := true
	curMax := 0
	step := 1
	for i := 0; i < len(aggNum); i += step {
		if skipLeadingNegetive == true {
			if aggNum[i] <= 0 {
				continue
			} else {
				skipLeadingNegetive = false
				step = 2
				curMax = aggNum[i]
			}

		} else {
			tmp := curMax + aggNum[i] + aggNum[i-1]
			if tmp > aggNum[i] {
				curMax = tmp
			} else {
				curMax = aggNum[i]
			}
		}
		if curMax > max {
			max = curMax
		}
	}

	return max

}

func maxSubArray5(nums []int) int {
	max := nums[0]
	maxEndingHere := max

	for i := 1; i < len(nums); i++ {
		if maxEndingHere+nums[i] > nums[i] {
			maxEndingHere = maxEndingHere + nums[i]
		} else {
			maxEndingHere = nums[i]
		}

		if maxEndingHere > max {
			max = maxEndingHere
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray1(nums))
	fmt.Println(maxSubArray2(nums))
	fmt.Println(maxSubArray3(nums))
	fmt.Println(maxSubArray4(nums))
	fmt.Println(maxSubArray5(nums))
}
