package problem

import (
	. "github.com/chrisxue815/leetcode-go/util"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"testing"
)

func convertBSTInorder(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	convertBSTInorder(root.Right, sum)

	*sum += root.Val
	root.Val = *sum

	convertBSTInorder(root.Left, sum)
}

// O(n) time. O(log(n)) space. In-order traversal.
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	convertBSTInorder(root, &sum)
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
		actual := convertBST(root)
		actualValues := actual.Serialize()
		assert.Equal(t, test.Expected, actualValues, "%+v", test)
	}
}
