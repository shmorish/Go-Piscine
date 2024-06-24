package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := BTreeLevelCount(root.Left)
	r := BTreeLevelCount(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
