package main

import (
	"fmt"
)

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else if data > root.Data {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}

	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}

	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else {
		return BTreeSearchItem(root.Right, elem)
	}
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}

func minValueNode(node *TreeNode) *TreeNode {
	current := node

	for current.Left != nil {
		current = current.Left
	}

	return current
}

func BTreeDeleteNode(root *TreeNode, key string) *TreeNode {
	if root == nil {
		return root
	}

	if key < root.Data {
		root.Left = BTreeDeleteNode(root.Left, key)
	} else if key > root.Data {
		root.Right = BTreeDeleteNode(root.Right, key)
	} else {

		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		temp := minValueNode(root.Right)

		root.Data = temp.Data

		root.Right = BTreeDeleteNode(root.Right, temp.Data)
	}

	return root
}

func main() {
	root := &TreeNode{Data: "4"}

	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	node := BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	BTreeApplyInorder(root, fmt.Println)
	root = BTreeDeleteNode(root, node.Data)
	fmt.Println("After delete:")
	BTreeApplyInorder(root, fmt.Println)
}
