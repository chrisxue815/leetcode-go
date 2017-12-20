package problem

import (
	. "leetcode-go/util"
)

func convertBSTInorder(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	convertBSTInorder(root.Right, sum)

	*sum += root.Val
	root.Val = *sum

	convertBSTInorder(root.Left, sum)
}

// O(n) time. O(log(n)) space. In-order traversal.
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	convertBSTInorder(root, &sum)
	return root
}
