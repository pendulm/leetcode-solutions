package main

import (
	"fmt"
	//  "math"
)

func Log2(N int) int {
	exp := 0
	for N != 0 {
		N = N >> 1
		exp += 1
	}
	return exp
}

func bitMask(width int) int {
	mask := ^0
	mask <<= uint(width)
	return ^mask

}

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	//  bitWidth := math.Ceil(math.Log2(float64(N)))
	bitWidth := Log2(N)
	//  mask := int(math.Exp2(bitWidth)) - 1
	mask := bitMask(bitWidth)
	return ^N & mask
}

func main() {
	n := 5
	fmt.Println(bitwiseComplement(n))
	n = 7
	fmt.Println(bitwiseComplement(n))
	n = 13
	fmt.Println(bitwiseComplement(n))
}
