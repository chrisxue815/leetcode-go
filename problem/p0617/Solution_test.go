package problem

import (
	. "github.com/chrisxue815/leetcode-go/util"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"testing"
)

// O(n) time. O(log(n)) space. Recursive post-order traversal.
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func Test_solution(t *testing.T) {
	type Args struct {
		T1 []null.Int `json:"t1"`
		T2 []null.Int `json:"t2"`
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
		t1 := Deserialize(test.Args.T1)
		t2 := Deserialize(test.Args.T2)
		actual := mergeTrees(t1, t2)
		actualValues := actual.Serialize()
		assert.Equal(t, test.Expected, actualValues, "%+v", test)
	}
}
