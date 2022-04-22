package problem

import (
	. "github.com/chrisxue815/leetcode-go/util"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"strconv"
	"testing"
)

func tree2strDFS(root *TreeNode, result []byte) []byte {
	result = strconv.AppendInt(result, int64(root.Val), 10)

	if root.Left != nil {
		result = append(result, '(')
		result = tree2strDFS(root.Left, result)
		result = append(result, ')')
	} else if root.Right != nil {
		result = append(result, "()"...)
	}

	if root.Right != nil {
		result = append(result, '(')
		result = tree2strDFS(root.Right, result)
		result = append(result, ')')
	}

	return result
}

// O(n) time. O(log(n)) space. Recursive pre-order traversal.
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	const initialCapacity = 16
	result := make([]byte, 0, initialCapacity)
	result = tree2strDFS(root, result)
	return string(result)
}

func Test_solution(t *testing.T) {
	type Args struct {
		Root []null.Int `json:"root"`
	}

	type TestCase struct {
		Args     Args   `json:"args"`
		Expected string `json:"expected"`
	}

	type TestJson struct {
		TestCases []TestCase `json:"test_cases"`
	}

	testJson := TestJson{}
	LoadTestJson(&testJson)

	for _, test := range testJson.TestCases {
		root := Deserialize(test.Args.Root)
		actual := tree2str(root)
		assert.Equal(t, test.Expected, actual, "%+v", test)
	}
}
