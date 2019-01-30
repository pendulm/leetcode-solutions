class Node:
    def __init__(self, value):
        self.value = value
        self.prev = None
        self.next = None


class DLinkList:
    def __init__(self):
        self.front = None
        self.tail = None

    def insert_front(self, node):
        old_front = self.front
        if old_front:
            node.prev = None
            node.next = old_front
        else:
            self.end = node
        self.front = node

    def pop_tail(self):
        if self.tail is None:
            return None

        if self.front is self.tail:
            node = self.front
            self.front = None
            self.end = None
            return node

        old_end = self.tail
        new_end = old_end.prev
        self.tail = new_end
        new_end.next = None



class LRUCache:

    def __init__(self, maxsize):
        self.maxsize = maxsize
        self._cache = {}
        self.dlist = DLinkList()

    def get(self, key):
        if key in self._cache:
            self._renew(key)
            return self._cache[key].value

        return None

    def _renew(self, key):


    def put(self, key, value):
        node = Node(value)
