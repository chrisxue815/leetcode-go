package problem

import (
	. "github.com/chrisxue815/leetcodego/util"
)

// O(n) time. O(log(n)) space. Recursive post-order traversal.
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
