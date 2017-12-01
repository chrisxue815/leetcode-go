package problem

import (
	. "leetcode-go/util"
)

// O(n) time. O(log(n)) space. Recursive DFS.
func trimBST(root *TreeNode, l int, r int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < l {
		return trimBST(root.Right, l, r)
	}
	if root.Val > r {
		return trimBST(root.Left, l, r)
	}

	root.Left = trimBST(root.Left, l, r)
	root.Right = trimBST(root.Right, l, r)
	return root
}
