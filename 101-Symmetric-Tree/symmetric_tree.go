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

func isFlippedTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}

	if left.Val != right.Val {
		return false
	}
	return isFlippedTree(left.Left, right.Right) && isFlippedTree(left.Right, right.Left)

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isFlippedTree(root.Left, root.Right)
}

func isSymmetricIter(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stackLeft := []*TreeNode{}
	stackRight := []*TreeNode{}

	stackLeft = append(stackLeft, root.Left)
	stackRight = append(stackRight, root.Right)
	var left *TreeNode = nil
	var right *TreeNode = nil

	for len(stackLeft) == len(stackRight) {
		if left == nil && right == nil {
			if len(stackLeft) == 0 {
				return true
			}
			left = stackLeft[len(stackLeft)-1]
			stackLeft = stackLeft[0 : len(stackLeft)-1]
			right = stackRight[len(stackRight)-1]
			stackRight = stackRight[0 : len(stackRight)-1]
			continue
		}

		if left == nil && right != nil || left != nil && right == nil {
			return false
		} else {
			if left.Val != right.Val {
				return false
			}
			stackLeft = append(stackLeft, left.Right)
			stackRight = append(stackRight, right.Left)
			left = left.Left
			right = right.Right
		}

	}
	return false
}

func main() {
	//  s := "1,2,2,3,4,4,3"
	//  s := "1,2,2,null,3,null,3"
	s := ""
	//  s := "1,2,2,3,4,4,3,5,6,7,8,8,7,6,5"
	arr := strings.Split(s, ",")
	tree := BuildTree(arr, 0)
	fmt.Println(isSymmetricIter(tree))
	//  fmt.Println(isSymmetric(tree))
}
