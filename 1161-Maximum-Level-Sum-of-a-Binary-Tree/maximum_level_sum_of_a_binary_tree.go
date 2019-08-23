package main

import (
	"container/list"
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

type LevelItem struct {
	node  *TreeNode
	level int
}

func dfs(root *TreeNode, levelSum map[int]int, level int) {
	if root == nil {
		return
	}

	levelSum[level] += root.Val
	dfs(root.Left, levelSum, level+1)
	dfs(root.Right, levelSum, level+1)
}

func bfs(root *TreeNode, levelSum map[int]int) {
	rootItem := &LevelItem{node: root, level: 1}
	queue := list.New()
	queue.PushBack(rootItem)

	for queue.Len() != 0 {
		item := queue.Remove(queue.Front()).(*LevelItem)
		levelSum[item.level] += item.node.Val
		if item.node.Left != nil {
			left := &LevelItem{node: item.node.Left, level: item.level + 1}
			queue.PushBack(left)
		}
		if item.node.Right != nil {
			right := &LevelItem{node: item.node.Right, level: item.level + 1}
			queue.PushBack(right)
		}
	}
}

func maxLevelSum(root *TreeNode) int {
	levelSum := make(map[int]int)

	dfs(root, levelSum, 1)
	//  bfs(root, levelSum)

	sum := root.Val
	maxLevel := 1
	for level := range levelSum {
		if levelSum[level] > sum {
			sum = levelSum[level]
			maxLevel = level
		}
	}
	return maxLevel
}

func main() {
	//  s := "1,7,0,7,-8,null,null"
	s := "63909,43838,4549,-31714,-99701,-96320,88666,75152,-14750,-12671,60405,null,29388,null,null,null,-83674,null,null,-83490,null,-49913,86188,84455,null,null,null,null,null,null,null,null,null,-36061,91438,-75550"
	arr := strings.Split(s, ",")
	tree := BuildTree(arr, 0)
	fmt.Println(maxLevelSum(tree))
}
