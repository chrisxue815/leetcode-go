package util

import (
	"strconv"
)

const initialCapacity = 16

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Serialize() string {
	buf := make([]byte, 0, initialCapacity)
	t.SerializeToBuffer(&buf)
	return string(buf)
}

func (t *TreeNode) SerializeToBuffer(buffer *[]byte) {
	if t == nil {
		return
	}

	buf := *buffer

	queue := make([]*TreeNode, 0, initialCapacity)
	queue = append(queue, t)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			buf = append(buf, "#,"...)
		} else {
			buf = strconv.AppendInt(buf, int64(node.Val), 10)
			buf = append(buf, ',')

			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	end := len(buf) - 1
	for ; end >= 0; end-- {
		last := buf[end]
		if last != '#' && last != ',' {
			break
		}
	}

	*buffer = buf[:end+1]
}

func Deserialize(values string) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	dummy := &TreeNode{}
	queue := make([]**TreeNode, 0, initialCapacity)
	queue = append(queue, &dummy)
	val := 0

	for i := 0; i < len(values); i++ {
		ch := values[i]

		switch ch {
		case ',':
			node := &TreeNode{
				Val: val,
			}
			*queue[0] = node
			queue = queue[1:]

			queue = append(queue, &node.Left)
			queue = append(queue, &node.Right)

			val = 0

		case '#':
			queue = queue[1:]
			i++

		default:
			val = val*10 + int(ch) - '0'
		}
	}

	*queue[0] = &TreeNode{
		Val: val,
	}

	return dummy
}
