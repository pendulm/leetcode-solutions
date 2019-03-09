package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := ListNode{0, nil}
    current :=  &head

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            current.Next = l1
            current = l1
            l1 = l1.Next
        } else {
            current.Next = l2
            current = l2
            l2 = l2.Next
        }
    }
    if (l1 == nil) {
        current.Next = l2
    } else {
        current.Next = l1
    }

    return head.Next
}

func BuildList(arr []int) *ListNode {
    if len(arr) == 0 {
        return nil
    }
    head := &ListNode{arr[0], nil}
    head.Next = BuildList(arr[1:])
    return head
}

func PrintList(head *ListNode) {
    for p := head; p != nil; p = p.Next {
        fmt.Printf("%d ", p.Val)
    }
    fmt.Printf("\n")
}

func main() {
    arr1 := []int{1,2,3,4}
    arr2 := []int{1,2,2}
    l1 := BuildList(arr1)
    PrintList(l1)
    l2 := BuildList(arr2)
    PrintList(l2)
    head := mergeTwoLists(l1, l2)
    PrintList(head)
}
