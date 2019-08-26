package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := 0
	if root.Left != nil {
		max = maxDepth(root.Left)
	}

	if root.Right != nil {
		rightMax := maxDepth(root.Right)
		if rightMax > max {
			max = rightMax
		}
	}

	return max + 1
}

func main() {
	s := "63909,43838,4549,-31714,-99701,-96320,88666,75152,-14750,-12671,60405,null,29388,null,null,null,-83674,null,null,-83490,null,-49913,86188,84455,null,null,null,null,null,null,null,null,null,-36061,91438,-75550"
	//  s := "3,9,20,null,null,15,7"
	arr := strings.Split(s, ",")
	tree := BuildTree(arr, 0)
	fmt.Println(maxDepth(tree))
}
