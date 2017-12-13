package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_addDigits(t *testing.T) {
	tests := []struct {
		num      int
		expected int
	}{
		{0, 0},
		{1, 1},
		{9, 9},
		{10, 1},
		{11, 2},
		{38, 2},
	}

	for _, test := range tests {
		actual := addDigits(test.num)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
