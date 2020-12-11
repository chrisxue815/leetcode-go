package problem

import (
	. "github.com/chrisxue815/leetcode-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

// O(n) time. O(1) space. Array, iteration.
func largeGroupPositions(s string) [][]int {
	result := make([][]int, 0)
	lo := 0

	for hi := 0; hi < len(s); hi++ {
		if hi == len(s)-1 || s[hi] != s[hi+1] {
			if hi-lo >= 2 {
				result = append(result, []int{lo, hi})
			}
			lo = hi + 1
		}
	}

	return result
}

func Test_solution(t *testing.T) {
	type Args struct {
		S string `json:"s"`
	}

	type TestCase struct {
		Args     Args `json:"args"`
		Expected [][]int  `json:"expected"`
	}

	type TestJson struct {
		TestCases []TestCase `json:"test_cases"`
	}

	testJson := TestJson{}
	LoadTestJson(&testJson)

	for _, test := range testJson.TestCases {
		actual := largeGroupPositions(test.Args.S)
		assert.Equal(t, test.Expected, actual, "%+v", test)
	}
}
