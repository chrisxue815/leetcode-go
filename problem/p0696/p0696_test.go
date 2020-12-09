package problem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countBinarySubstrings(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"00110011", 6},
		{"10101", 4},
	}

	for _, test := range tests {
		actual := countBinarySubstrings(test.s)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
