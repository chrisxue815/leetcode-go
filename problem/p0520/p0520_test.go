package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_detectCapitalUse(t *testing.T) {
	tests := []struct {
		word     string
		expected bool
	}{
		{"USA", true},
		{"leetcode", true},
		{"Google", true},
		{"FlaG", false},
		{"uSA", false},
	}

	for _, test := range tests {
		actual := detectCapitalUse(test.word)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
