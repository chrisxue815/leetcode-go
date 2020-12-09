package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_hammingDistance(t *testing.T) {
	tests := []struct {
		x        int
		y        int
		expected int
	}{
		{1, 4, 2},
	}

	for _, test := range tests {
		actual := hammingDistance(test.x, test.y)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
