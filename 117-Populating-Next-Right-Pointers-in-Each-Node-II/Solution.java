class Solution {

    // root is not null
    void linkNextOrParent(Node root) {
        if (root.left != null) {
            if (root.right != null) {
                root.left.next = root.right;
            } else {
                root.left.next = root;
            }
            linkNextOrParent(root.left);
        }
        if (root.right != null) {
            root.right.next = root;
            linkNextOrParent(root.right);
        }
    }

    boolean nextToParent(Node node) {
        Node next = node.next;
        if (next == null) {
            return false;
        }

        if (next.left == node || next.right == node) {
            return true;
        }
        return false;
    }

    Node nextRight(Node node) {
        Node parent = null;
        if (nextToParent(node) == true) {
            parent = node.next;
            while (parent.next != null) {
                parent = parent.next;
                if (parent.left != null) {
                    return parent.left;
                }
                if (parent.right != null) {
                    return parent.right;
                }
            }
            return null;
        }
        return node.next;
    }


    public Node connect(Node root) {
        if (root == null) {
            return null;
        }

        linkNextOrParent(root);

        Node leftMost = root;
        Node current = null;
        Node next = null;

        while (leftMost != null) {
            current = leftMost;
            leftMost = null;
            do {
                next = nextRight(current);
                current.next = next;
                if (leftMost == null) {
                    if (current.left != null) {
                        leftMost = current.left;
                    } else if (current.right != null) {
                        leftMost = current.right;
                    }
                }
                current = next;
            } while(current != null);
        }

        return root;
        
    }

    void displayTree(Node node) {
        if (node == null) {
            return;
        }
        if (node.next == null) {
            System.out.printf("%d->%s\n", node.val, "nil");
        } else {
            System.out.printf("%d->%d\n", node.val, node.next.val);
        }

        displayTree(node.left);
        displayTree(node.right);


    }

    public static void main(String[] args) {
        Node node4 = new Node(4, null, null, null);
        Node node5 = new Node(5, null, null, null);
        Node node7 = new Node(7, null, null, null);
        Node node2 = new Node(2, node4, node5, null);
        Node node3 = new Node(3, null, node7, null);
        Node node1 = new Node(1, node2, node3, null);
        Solution s = new Solution();
        Node root = s.connect(node1);
        s.displayTree(root);
    }
}
