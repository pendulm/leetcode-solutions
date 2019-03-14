package main

import (
    "fmt"
)

func reverse(x int) int {
    y := x
    if (x < 0 ) {
        y = -x
    }

    reverseY := y % 10
    y = y / 10
    for y != 0 {
        reverseY = reverseY * 10 + y % 10
        y = y / 10
    }

    if (reverseY >= 1<<31) {
        return 0
    }


    if (x < 0) {
        reverseY = -reverseY
    }
    return reverseY

}


func main() {
    n := 123
    fmt.Println(reverse(n))
    n = -123
    fmt.Println(reverse(n))
    n = 120
    fmt.Println(reverse(n))
    n = 1534236469
    fmt.Println(reverse(n))
}
