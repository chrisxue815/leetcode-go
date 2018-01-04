package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_findRelativeRanks(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []string
	}{
		{[]int{5, 4, 3, 2, 1}, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
	}

	for _, test := range tests {
		actual := findRelativeRanks(test.nums)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
