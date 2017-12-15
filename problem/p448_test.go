package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_findDisappearedNumbers(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
	}

	for _, test := range tests {
		actual := findDisappearedNumbers(test.nums)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
