package test_0136

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		n        []int
		expected int
	}{
		{[]int{1, 2, 1}, 2},
	}

	for _, test := range tests {
		actual := singleNumber(test.n)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
