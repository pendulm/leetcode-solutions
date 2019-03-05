package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if (head == nil || k == 0) {
        return head
    }

    length := 1
    p := head
    for  p.Next != nil {
        p = p.Next
        length += 1
    }

    k = k % length
    if (k == 0) {
        return head
    }
    p.Next = head

    k = length - k

    p = head
    for i := 0; i < k - 1; i++ {
        p = p.Next
    }
    head = p.Next
    p.Next = nil
    return head
}

func PrintList(head *ListNode) {
    for p := head; p != nil; p = p.Next {
        fmt.Printf("%d -> ", p.Val)
    }
    fmt.Println("nil")
}

func BuildList(seq []int) *ListNode {
    if (len(seq) == 0) {
        return nil
    }
    node := &ListNode{seq[0], nil}
    node.Next = BuildList(seq[1:])
    return node
}

func main() {
    list := BuildList([]int{})
    //  list := BuildList([]int{1,2,3,4,5})
    PrintList(list)
    list = rotateRight(list, 11)
    PrintList(list)
}
