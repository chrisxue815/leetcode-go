package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"hello", "olleh"},
	}

	for _, test := range tests {
		actual := reverseString(test.s)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
