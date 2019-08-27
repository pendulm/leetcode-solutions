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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	leftLen := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			leftLen = i
		}
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(inorder[0:leftLen], postorder[0:leftLen]),
		Right: buildTree(inorder[leftLen+1:], postorder[leftLen:len(postorder)-1]),
	}
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
	LevelOrderPrint(root)
}
