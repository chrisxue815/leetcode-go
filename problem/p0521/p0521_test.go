package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_findLUSlength(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected int
	}{
		{"aba", "cdc", 3},
		{"abc", "abc", -1},
		{"abcd", "abc", 4},
		{"abc", "abcd", 4},
	}

	for _, test := range tests {
		actual := findLUSlength(test.a, test.b)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
