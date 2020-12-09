package problem

import (
	. "github.com/chrisxue815/leetcode-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxDistToClosest(seats []int) int {
	lo := 0

	for lo < len(seats) && seats[lo] == 0 {
		lo++
	}

	result := lo

	for hi := lo + 1; hi < len(seats); hi++ {
		if seats[hi] == 1 {
			dist := (hi - lo) / 2
			if result < dist {
				result = dist
			}
			lo = hi
		}
	}

	dist := len(seats) - 1 - lo
	if result < dist {
		result = dist
	}

	return result
}

func Test_solution(t *testing.T) {
	type Args struct {
		Seats []int `json:"seats"`
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
		actual := maxDistToClosest(test.Args.Seats)
		assert.Equal(t, test.Expected, actual, "%+v", test)
	}
}
