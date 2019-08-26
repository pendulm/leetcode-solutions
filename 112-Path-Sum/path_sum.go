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

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	lOk := false
	rOk := false

	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return true
		}
	}

	if root.Left != nil {
		lOk = hasPathSum(root.Left, sum-root.Val)
	}
	if root.Right != nil {
		rOk = hasPathSum(root.Right, sum-root.Val)
	}

	return lOk || rOk

}

func main() {
	s := "1,2"
	//  s := "5,4,8,11,null,13,4,7,2,null,null,null,null,null,1"
	arr := strings.Split(s, ",")
	tree := BuildTree(arr, 0)
	fmt.Println(hasPathSum(tree, 1))
	//  fmt.Println(isSymmetric(tree))
}
