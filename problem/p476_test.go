package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindComplement(t *testing.T) {
	tests := []struct {
		num      int
		expected int
	}{
		{5, 2},
		{1, 0},
	}

	for _, test := range tests {
		actual := findComplement(test.num)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
