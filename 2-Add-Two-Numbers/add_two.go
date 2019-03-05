package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func listToInt(l *ListNode) int {
    var num int = 0
    var base int = 1
    for l != nil {
        num = num + l.Val*base
        l = l.Next
        base = base*10
    }
    return num
}

func intToList(num int) *ListNode {
    if (num < 10) {
        return &ListNode{num, nil}
    }

    rem := num % 10
    tail := intToList(num / 10)
    head := &ListNode{rem, tail}
    return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    var head *ListNode = nil
    var tail *ListNode = nil
    carry := 0
    digit := 0

    for  {
        if (l1 == nil && l2 == nil) {
            break
        } else {
            if (l1 != nil && l2 == nil) {
                digit, carry = digitAdd(l1.Val, 0, carry)
            } else if (l1 == nil && l2 != nil) {
                digit, carry = digitAdd(0, l2.Val, carry)
            } else {
                digit, carry = digitAdd(l1.Val, l2.Val, carry)
            }

            if (tail == nil) {
                head = &ListNode{digit, nil}
                tail = head
            } else {
                tail.Next = &ListNode{digit, nil}
                tail = tail.Next
            }

            if (l1 != nil) {
                l1 = l1.Next
            }
            if (l2 != nil) {
                l2 = l2.Next
            }
        }
    }

    if (carry == 1) {
        tail.Next = &ListNode{1, nil}
    }

    return head
}

func digitAdd(a int, b int, carry int) (int, int) {
    addNum := a + b + carry
    return addNum % 10, addNum / 10
}

func printList(l *ListNode) {
    for l != nil {
        fmt.Printf("%d", l.Val)
        l = l.Next
    }
    fmt.Printf("\n")
}

func main() {
    a := 9
    b := 9
    aL := intToList(a)
    printList(aL)
    bL := intToList(b)
    printList(bL)
    sumL := addTwoNumbers(aL, bL)
    fmt.Printf("expect: %d, result: %d\n", a+b, listToInt(sumL))
}
