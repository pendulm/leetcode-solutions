package main

import (
    "fmt"
)

func decidePalindrome(s string) bool {
    for left, right := 0, len(s)-1; left <= right; left, right = left+1, right-1 {
        if s[left] != s[right] {
            return false
        }
    }
    return true
}

func longestPalindrome(s string) string {
    length := len(s)
    if (length == 0) {
        return ""
    }
    for curLen := length; curLen > 0; curLen-- {
        for start := 0; start <= length - curLen; start++  {
            subStr := s[start:start+curLen]
            if decidePalindrome(subStr) == true {
                return subStr
            }
        }
    }
    return s[0:1]
}


func main() {
    s := "Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 10000."
    fmt.Println(longestPalindrome(s))
}
