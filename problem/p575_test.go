package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDistributeCandies(t *testing.T) {
	tests := []struct {
		candies  []int
		expected int
	}{
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{1, 1, 2, 3}, 2},
		{[]int{1, 1, 1, 2, 2, 2, 3, 3, 3}, 3},
	}

	for _, test := range tests {
		actual := distributeCandies(test.candies)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
