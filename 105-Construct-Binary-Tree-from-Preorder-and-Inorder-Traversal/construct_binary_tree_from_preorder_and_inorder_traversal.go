package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrderPrint(root *TreeNode) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		levelLen := queue.Len()
		for i := 0; i < levelLen; i++ {
			item := queue.Remove(queue.Front()).(*TreeNode)
			if item.Left != nil {
				queue.PushBack(item.Left)
			}
			if item.Right != nil {
				queue.PushBack(item.Right)
			}
			fmt.Printf("%d ", item.Val)
		}
		fmt.Printf("\n")
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	leftLen := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			leftLen = i
		}
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:leftLen+1], inorder[0:leftLen]),
		Right: buildTree(preorder[leftLen+1:], inorder[leftLen+1:]),
	}
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	LevelOrderPrint(root)
}
