package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_judgeCircle(t *testing.T) {
	tests := []struct {
		moves    string
		expected bool
	}{
		{"UD", true},
		{"LL", false},
	}

	for _, test := range tests {
		actual := judgeCircle(test.moves)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
