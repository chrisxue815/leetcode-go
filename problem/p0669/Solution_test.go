package problem

import (
	. "github.com/chrisxue815/leetcode-go/util"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"testing"
)

// O(n) time. O(log(n)) space. Recursive post-order traversal.
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

func Test_solution(t *testing.T) {
	type Args struct {
		Root []null.Int `json:"root"`
		Low  int        `json:"low"`
		High int        `json:"high"`
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
		actual := trimBST(root, test.Args.Low, test.Args.High)
		actualValues := actual.Serialize()
		assert.Equal(t, test.Expected, actualValues, "%+v", test)
	}
}
