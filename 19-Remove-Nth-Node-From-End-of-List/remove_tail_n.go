package main

import (
    "fmt"
)


type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if (head == nil) {
        return head
    }

    p1 := head
    p2 := head
    for i:= 0; i < n; i++ {
        p2 = p2.Next
    }

    if p2 == nil {
        return head.Next
    }

    for p2.Next != nil {
        p1 = p1.Next
        p2 = p2.Next
    }
    p1.Next = p1.Next.Next

    return head
}

func buildList(arr []int) *ListNode {
    if (len(arr) == 0) {
        return nil
    }

    head := &ListNode{arr[0], nil}
    head.Next = buildList(arr[1:])
    return head
}

func printList(head *ListNode) {
    for head != nil {
        fmt.Printf("%d ", head.Val)
        head = head.Next
    }
    fmt.Printf("\n")
}

func main() {
    arr := []int{1}
    //  arr := []int{1,2,3,4,5}
    head := buildList(arr)
    list := removeNthFromEnd(head, 1)
    printList(list)
}
