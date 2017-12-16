package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_getSum(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{0, 2, 2},
		{2, 0, 2},
		{1, -2, -1},
		{2, -1, 1},
	}

	for _, test := range tests {
		actual := getSum(test.a, test.b)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
