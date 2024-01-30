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

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		f(root.Data)
		BTreeApplyPreorder(root.Left, f)
		BTreeApplyPreorder(root.Right, f)
	}
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
