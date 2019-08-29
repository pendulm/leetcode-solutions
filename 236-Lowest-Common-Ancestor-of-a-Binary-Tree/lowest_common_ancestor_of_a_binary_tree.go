package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(arr []string, index int) *TreeNode {
	if index >= len(arr) {
		return nil
	}

	i, err := strconv.Atoi(arr[index])
	if err != nil {
		return nil
	}

	left := BuildTree(arr, index*2+1)
	right := BuildTree(arr, index*2+2)
	return &TreeNode{
		i,
		left,
		right,
	}
}

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	t := findNode(root.Left, val)
	if t != nil {
		return t
	}
	t = findNode(root.Right, val)
	if t != nil {
		return t
	}
	return nil
}

func inSubtree(root, node *TreeNode) bool {
	if root == nil {
		return false
	}
	return (root == node) || inSubtree(root.Left, node) || inSubtree(root.Right, node)
}

func doSearch(current, p, q *TreeNode) (bool, *TreeNode) {
	if current == nil {
		return false, nil
	}
	// q cant be in subtree of p
	if current == p {
		return true, nil
	}
	findP, target := doSearch(current.Left, p, q)
	if target != nil {
		return true, target
	} else if findP == true {
		if inSubtree(current, q) {
			return true, current
		} else {
			return true, nil
		}
	}
	findP, target = doSearch(current.Right, p, q)
	if target != nil {
		return true, target
	} else if findP == true {
		if inSubtree(current, q) {
			return true, current
		} else {
			return true, nil
		}
	}

	return false, nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if inSubtree(p, q) {
		return p
	} else if inSubtree(q, p) {
		return q
	}
	_, target := doSearch(root, p, q)
	return target
}

func testCase(treeRepr string, p int, q int, result int) bool {
	part := strings.Split(treeRepr, ",")
	root := BuildTree(part, 0)
	nodeP := findNode(root, p)
	nodeQ := findNode(root, q)
	nodeCommon := findNode(root, result)
	if nodeCommon == lowestCommonAncestor(root, nodeP, nodeQ) {
		return true

	}
	return false
}

func main() {
	fmt.Println(testCase("3,5,1,6,2,0,8,null,null,7,4", 5, 1, 3))
	fmt.Println(testCase("3,5,1,6,2,0,8,null,null,7,4", 5, 4, 5))
}
