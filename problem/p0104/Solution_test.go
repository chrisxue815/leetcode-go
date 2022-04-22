package problem

import (
	. "github.com/chrisxue815/leetcode-go/util"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"testing"
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

func Test_solution(t *testing.T) {
	type Args struct {
		Root []null.Int `json:"root"`
	}

	type TestCase struct {
		Args     Args `json:"args"`
		Expected int  `json:"expected"`
	}

	type TestJson struct {
		TestCases []TestCase `json:"test_cases"`
	}

	testJson := TestJson{}
	LoadTestJson(&testJson)

	for _, test := range testJson.TestCases {
		root := Deserialize(test.Args.Root)
		actual := maxDepth(root)
		assert.Equal(t, test.Expected, actual, "%+v", test)
	}
}
