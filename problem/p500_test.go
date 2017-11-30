package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindWords(t *testing.T) {
	tests := []struct {
		words    []string
		expected []string
	}{
		{[]string{"Hello", "Alaska", "Dad", "Peace"}, []string{"Alaska", "Dad"}},
	}

	for _, test := range tests {
		actual := findWords(test.words)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
