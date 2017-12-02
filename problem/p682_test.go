package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalPoints(t *testing.T) {
	tests := []struct {
		ops      []string
		expected int
	}{
		{[]string{"5", "2", "C", "D", "+"}, 30},
		{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
	}

	for _, test := range tests {
		actual := calPoints(test.ops)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
