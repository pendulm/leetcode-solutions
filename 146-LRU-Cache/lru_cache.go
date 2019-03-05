package main

import (
    "fmt"
)

type DLinkNode struct {
    key int
    value int
    prev *DLinkNode
    next *DLinkNode
}

type LRUCache struct {
    hash map[int]*DLinkNode
    head *DLinkNode
    tail *DLinkNode
    capacity int
    size int
}


func Constructor(capacity int) LRUCache {
    lru := new(LRUCache)
    lru.capacity = capacity
    lru.hash = make(map[int]*DLinkNode)
    return *lru
}


func (this *LRUCache) Get(key int) int {
    node, ok := this.hash[key]
    if !ok {
        return -1
    }

    if (node != this.head) {
        node.prev.next = node.next
        if (node.next != nil) {
            node.next.prev = node.prev
        }
        if (node == this.tail) {
            this.tail = node.prev
        }
        node.next = this.head
        this.head.prev = node
        this.head = node
    }
    return node.value

}


func (this *LRUCache) Put(key int, value int)  {
    var node *DLinkNode

    if (this.Get(key) != -1) {
        node = this.hash[key]
        node.value = value
        return
    }

    node = &DLinkNode{key, value, nil, nil}
    node.next = this.head
    if (this.head != nil) {
        this.head.prev = node
    }
    if (this.tail == nil) {
        this.tail = node
    }
    if (this.size < this.capacity) {
        this.size += 1
    } else {
        delete(this.hash, this.tail.key)
        this.tail.prev.next = nil
        this.tail = this.tail.prev
    }

    this.hash[key] = node
    this.head = node
}

func (this *LRUCache) PrintList() {
    for node := this.head; node != nil; node = node.next {
        fmt.Printf("%d -> ", node.value)
    }
    fmt.Println("nil")
}


func main() {
    obj := Constructor(5);
    obj.Put(1, 1)
    obj.PrintList()
    obj.Put(2, 2)
    obj.PrintList()
    obj.Put(3, 3)
    obj.PrintList()
    obj.Put(4, 4)
    obj.PrintList()
    obj.Put(4, 5)
    obj.PrintList()
    obj.Put(4, 6)
    obj.PrintList()
    obj.Put(4, 7)
    obj.PrintList()
    obj.Get(2)
    obj.PrintList()
    obj.Get(3)
    obj.PrintList()
}
