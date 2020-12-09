package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}

	for _, test := range tests {
		moveZeroes(test.nums)
		assert.Equal(t, test.expected, test.nums, "%+v", test)
	}
}
