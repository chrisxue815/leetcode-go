package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_matrixReshape(t *testing.T) {
	tests := []struct {
		nums     [][]int
		rows     int
		cols     int
		expected [][]int
	}{
		{
			[][]int{
				{1, 2},
				{3, 4},
			},
			1,
			4,
			[][]int{
				{1, 2, 3, 4},
			},
		},
		{
			[][]int{
				{1, 2},
				{3, 4},
			},
			2,
			4,
			[][]int{
				{1, 2},
				{3, 4},
			},
		},
	}

	for _, test := range tests {
		actual := matrixReshape(test.nums, test.rows, test.cols)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
