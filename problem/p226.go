package problem

import (
	. "leetcode-go/util"
)

// O(n) time. O(log(n)) space. Recursive pre-order traversal.
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
