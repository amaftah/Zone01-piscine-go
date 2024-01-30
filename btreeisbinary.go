package piscine

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

func BTreeIsBinary(root *TreeNode) bool {
	return isBSTUtil(root, "", "")
}

func isBSTUtil(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}

	if (min != "" && node.Data <= min) || (max != "" && node.Data >= max) {
		return false
	}

	return isBSTUtil(node.Left, min, node.Data) && isBSTUtil(node.Right, node.Data, max)
}
