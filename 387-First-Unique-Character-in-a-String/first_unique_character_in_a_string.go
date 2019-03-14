package main

import (
    "fmt"
)

func firstUniqChar(s string) int {
    var charNum [26]int
    for i := 0; i < len(s); i++ {
        charIndex := s[i] - 'a'
        if charNum[charIndex] == 0 {
            charNum[charIndex] = i+1
        } else if charNum[charIndex] > 0 {
            charNum[charIndex] = -1
        }
    }

    firstUniqPos := len(s) + 1
    for _, pos := range(charNum) {
        if pos > 0 && pos < firstUniqPos {
            firstUniqPos = pos
        }
    }
    if (firstUniqPos == len(s)+1) {
        return -1
    } else {
        return firstUniqPos - 1
    }
}

func main() {
    s := "leetcode"
    fmt.Println(firstUniqChar(s))
    s = "loveleetcode"
    fmt.Println(firstUniqChar(s))
    s = "aaaaaaaa"
    fmt.Println(firstUniqChar(s))
    s = "z"
    fmt.Println(firstUniqChar(s))
}
