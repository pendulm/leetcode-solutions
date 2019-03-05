package main

import (
    "fmt"
)

func lengthOfLongestSubstring(s string) int {
    seenChar := make(map[byte]int)
    maxLen := 0
    curLen := 0

    for i := 0; i < len(s); i++ {
        c := s[i]
        lastIndex, ok := seenChar[c]
        if !ok {
            curLen += 1
        } else {
            if (curLen > maxLen) {
                maxLen = curLen
            }
            if (i - curLen < lastIndex) {
                curLen = i - lastIndex
            }
        }
        seenChar[c] = i
    }
    if (curLen > maxLen) {
        maxLen = curLen
    }

    return maxLen
}

func main() {
    input := "abcabcbb"
    fmt.Println(lengthOfLongestSubstring(input))
    input = "bbbbb"
    fmt.Println(lengthOfLongestSubstring(input))
    input = "pwwkew"
    fmt.Println(lengthOfLongestSubstring(input))
    input = "abba"
    fmt.Println(lengthOfLongestSubstring(input))
}
