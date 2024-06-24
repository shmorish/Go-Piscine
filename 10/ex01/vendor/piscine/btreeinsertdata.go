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
		if root.Left == nil {
			root.Left = &TreeNode{Data: data}
			root.Left.Parent = root
			return root.Left
		}
		return BTreeInsertData(root.Left, data)
	}
	if data > root.Data {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data}
			root.Right.Parent = root
			return root.Right
		}
		return BTreeInsertData(root.Right, data)
	}
	return root
}
