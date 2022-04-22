package problem

import (
	"container/list"
	. "github.com/chrisxue815/leetcode-go/util"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"testing"
)

// O(n) time. O(n) space. BFS.
func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		result = append(result, queue.Back().Value.(*TreeNode).Val)
		length := queue.Len()

		for i := 0; i < length; i++ {
			curr := queue.Remove(queue.Front()).(*TreeNode)
			if curr.Left != nil {
				queue.PushBack(curr.Left)
			}
			if curr.Right != nil {
				queue.PushBack(curr.Right)
			}
		}
	}

	return result
}

func Test_solution(t *testing.T) {
	type Args struct {
		Root []null.Int `json:"root"`
	}

	type TestCase struct {
		Args     Args  `json:"args"`
		Expected []int `json:"expected"`
	}

	type TestJson struct {
		TestCases []TestCase `json:"test_cases"`
	}

	testJson := TestJson{}
	LoadTestJson(&testJson)

	for _, test := range testJson.TestCases {
		root := Deserialize(test.Args.Root)
		actual := rightSideView(root)
		assert.Equal(t, test.Expected, actual, "%+v", test)
	}
}
