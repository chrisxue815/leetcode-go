package problem

import (
	. "github.com/chrisxue815/leetcode-go/util"
)

import (
	"strconv"
)

func tree2strDFS(root *TreeNode, result []byte) []byte {
	result = strconv.AppendInt(result, int64(root.Val), 10)

	if root.Left != nil {
		result = append(result, '(')
		result = tree2strDFS(root.Left, result)
		result = append(result, ')')
	} else if root.Right != nil {
		result = append(result, "()"...)
	}

	if root.Right != nil {
		result = append(result, '(')
		result = tree2strDFS(root.Right, result)
		result = append(result, ')')
	}

	return result
}

// O(n) time. O(log(n)) space. Recursive pre-order traversal.
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	const initialCapacity = 16
	result := make([]byte, 0, initialCapacity)
	result = tree2strDFS(root, result)
	return string(result)
}
