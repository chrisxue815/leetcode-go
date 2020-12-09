package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_nextGreaterElement(t *testing.T) {
	tests := []struct {
		findNums []int
		nums     []int
		expected []int
	}{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{[]int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
	}

	for _, test := range tests {
		actual := nextGreaterElement(test.findNums, test.nums)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
