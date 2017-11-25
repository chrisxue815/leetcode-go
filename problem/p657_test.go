package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestJudgeCircle(t *testing.T) {
	tests := []struct {
		moves    string
		expected bool
	}{
		{"UD", true},
		{"LL", false},
	}

	for _, test := range tests {
		actual := judgeCircle(test.moves)
		assert.Equal(t, test.expected, actual, "moves=%s", test.moves)
	}
}
