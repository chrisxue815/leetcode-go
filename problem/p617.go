package problem

import (
	. "leetcode-go/util"
)

// O(n) time. O(log(n)) space. Recursive pre-order traversal.
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 != nil {
		if t2 != nil {
			t1.Val += t2.Val
			t1.Left = mergeTrees(t1.Left, t2.Left)
			t1.Right = mergeTrees(t1.Right, t2.Right)
			return t1
		} else {
			return t1
		}
	} else {
		return t2
	}
}
