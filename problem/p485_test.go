package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
	}

	for _, test := range tests {
		actual := findMaxConsecutiveOnes(test.nums)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
