package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_findTheDifference(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected byte
	}{
		{"abcd", "abcde", 'e'},
	}

	for _, test := range tests {
		actual := findTheDifference(test.s, test.t)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
