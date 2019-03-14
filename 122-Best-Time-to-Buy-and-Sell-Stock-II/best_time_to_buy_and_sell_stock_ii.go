package main

import (
    "fmt"
)

func maxProfit(prices []int) int {
    profit := 0
    risePos := 0

    length := len(prices)

    for i := 0; i < length-1; i++  {
        if prices[i] > prices[i+1] {
            profit += prices[i] - prices[risePos]
            risePos = i+1
        } else {
            if (i == length-2) {
                profit += prices[length-1] - prices[risePos]
            }
        }
    }

    return profit
}


func main() {
    prices := []int{7,1,5,3,6,4}
    fmt.Println(maxProfit(prices))
    prices = []int{1,2,3,4,5}
    fmt.Println(maxProfit(prices))
    prices = []int{7,6,4,3,1}
    fmt.Println(maxProfit(prices))
}
