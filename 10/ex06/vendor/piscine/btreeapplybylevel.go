package piscine

func QueueLen(queue []*TreeNode) int {
	count := 0
	for range queue {
		count++
	}
	return count
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for QueueLen(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		f(node.Data)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
