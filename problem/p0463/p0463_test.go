package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_islandPerimeter(t *testing.T) {
	tests := []struct {
		grid     [][]int
		expected int
	}{
		{
			[][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			16,
		},
	}

	for _, test := range tests {
		actual := islandPerimeter(test.grid)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
