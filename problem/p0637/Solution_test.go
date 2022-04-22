package problem

import (
	. "github.com/chrisxue815/leetcode-go/util"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"testing"
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

func Test_solution(t *testing.T) {
	type Args struct {
		Root []null.Int `json:"root"`
	}

	type TestCase struct {
		Args     Args      `json:"args"`
		Expected []float64 `json:"expected"`
	}

	type TestJson struct {
		TestCases []TestCase `json:"test_cases"`
	}

	testJson := TestJson{}
	LoadTestJson(&testJson)

	for _, test := range testJson.TestCases {
		root := Deserialize(test.Args.Root)
		actual := averageOfLevels(root)
		assert.Equal(t, test.Expected, actual, "%+v", test)
	}
}
