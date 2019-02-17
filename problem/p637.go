package problem

import (
	. "github.com/chrisxue815/leetcodego/util"
)

// O(n) time. O(n) space. BFS.
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	const initialCapacity = 16
	result := make([]float64, 0, initialCapacity)
	queue := make([]*TreeNode, 0, initialCapacity)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelLength := len(queue)
		sum := 0

		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]

			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, float64(sum)/float64(levelLength))
	}

	return result
}
