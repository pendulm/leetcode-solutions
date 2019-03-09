package main

import (
    "fmt"
)

func isValid(s string) bool {
    if (len(s) == 0) {
        return true
    }
    stack := []rune{}
    stackLen := 0

    for _, b := range(s) {
        switch b {
        case '[', '(', '{':
            stack = append(stack, b)
            stackLen += 1
        case ']', ')', '}':
            if (stackLen == 0) {
                return false
            }
            toMatch := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            stackLen -= 1
            if b == ']' && toMatch != '[' {
                return false
            } else if b == ')' && toMatch != '(' {
                return false
            } else if b == '}' && toMatch != '{' {
                return false
            }
        default:
            return false
        }

    }

    if (stackLen != 0) {
        return false
    }

    return true
}

func main() {
    input := ""
    fmt.Println(input, isValid(input))
    input = "()"
    fmt.Println(input, isValid(input))
    input = "()[]{}"
    fmt.Println(input, isValid(input))
    input = "(]"
    fmt.Println(input, isValid(input))
    input = "([)]"
    fmt.Println(input, isValid(input))
    input = "{[]}"
    fmt.Println(input, isValid(input))
    input = "{[]"
    fmt.Println(input, isValid(input))
}
