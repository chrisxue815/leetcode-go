package problem

import (
	. "github.com/chrisxue815/leetcode-go/util"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"testing"
)

// O(n) time. O(log(n)) space. Recursive pre-order traversal.
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func Test_solution(t *testing.T) {
	type Args struct {
		Root []null.Int `json:"root"`
	}

	type TestCase struct {
		Args     Args       `json:"args"`
		Expected []null.Int `json:"expected"`
	}

	type TestJson struct {
		TestCases []TestCase `json:"test_cases"`
	}

	testJson := TestJson{}
	LoadTestJson(&testJson)

	for _, test := range testJson.TestCases {
		root := Deserialize(test.Args.Root)
		actual := invertTree(root)
		actualValues := actual.Serialize()
		assert.Equal(t, test.Expected, actualValues, "%+v", test)
	}
}
