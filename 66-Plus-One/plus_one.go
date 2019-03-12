package main

import (
    "fmt"
)

func plusOne(digits []int) []int {
    carry := 1
    length := len(digits)
    for i := len(digits) - 1; i >= 0; i-- {
        sum := digits[i] + carry
        carry = sum / 10
        digits[i] = sum % 10

        if (carry == 0) {
            break
        }
    }
    if (carry == 1) {
        digits = append(digits, 0)
        copy(digits[0:length], digits[1:length+1])
        digits[0] = 1
    }
    return digits
}


func main() {
    digits := []int{9,9,9,9}
    digits = plusOne(digits)
    fmt.Println(digits)
    digits = []int{1}
    digits = plusOne(digits)
    fmt.Println(digits)
}
