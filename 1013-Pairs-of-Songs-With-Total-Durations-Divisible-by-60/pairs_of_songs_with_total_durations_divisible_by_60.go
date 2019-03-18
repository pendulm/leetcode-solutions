package main

import (
	"fmt"
)

//  func numPairsDivisibleBy60(time []int) int {
//  count := 0
//  length := len(time)
//  if length < 2 {
//  return 0
//  }

//  for i := 0; i < length-1; i++ {
//  for j := i + 1; j < length; j++ {
//  if (time[i]+time[j])%60 == 0 {
//  count += 1
//  }
//  }
//  }
//  return count
//  }

func numPairsDivisibleBy60(time []int) int {
	complete := [60]int{}
	count := 0

	for _, t := range time {
		count += complete[t%60]
		complete[(60-t%60)%60] += 1
	}
	return count
}

func main() {
	arr := []int{30, 20, 150, 100, 40}
	fmt.Println(numPairsDivisibleBy60(arr))
	arr = []int{60, 60, 60}
	fmt.Println(numPairsDivisibleBy60(arr))
}
