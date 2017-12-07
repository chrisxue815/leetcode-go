package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_hasAlternatingBits(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{1, true},
		{2, true},
		{5, true},
		{7, false},
		{10, true},
		{11, false},
	}

	for _, test := range tests {
		actual := hasAlternatingBits(test.n)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
