package util

import (
	"gopkg.in/guregu/null.v4"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Serialize() []null.Int {
	result := make([]null.Int, 0)
	if t == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, t)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, null.NewInt(0, false))
		} else {
			result = append(result, null.NewInt(int64(node.Val), true))

			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	for i := len(result) - 1; i >= 0; i-- {
		if result[i].Valid {
			break
		}
		result = result[:len(result)-1]
	}

	return result
}

func Deserialize(values []null.Int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	var root *TreeNode = nil
	queue := make([]**TreeNode, 0)
	queue = append(queue, &root)

	for i := 0; i < len(values); i++ {
		value := values[i]

		if value.Valid {
			node := &TreeNode{
				Val: int(value.Int64),
			}
			*queue[0] = node
			queue = queue[1:]

			queue = append(queue, &node.Left)
			queue = append(queue, &node.Right)
		} else {
			queue = queue[1:]
		}
	}

	return root
}
