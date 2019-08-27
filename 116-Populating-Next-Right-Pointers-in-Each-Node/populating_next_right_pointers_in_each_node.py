#!/usr/bin/env python3

import queue

class Node:
    def __init__(self, val, left, right, next):
        self.val = val
        self.left = left
        self.right = right
        self.next = next

class Solution:
    def connect(self, root: 'Node') -> 'Node':
        if root is None:
            return None

        q = queue.Queue()
        q.put(root)

        while not q.empty():
            qlen = q.qsize()
            for i in range(qlen):
                item = q.get()
                item.next = None
                if item.left != None:
                    q.put(item.left)
                if item.right != None:
                    q.put(item.right)

                if i == 0:
                    prev = item
                else:
                    prev.next = item
                    prev = item
        return root

