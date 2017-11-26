package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestArrayPairSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2}, 1},
		{[]int{1, 4, 3, 2}, 4},
	}

	for _, test := range tests {
		actual := arrayPairSum(test.nums)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
