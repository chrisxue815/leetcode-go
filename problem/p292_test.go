package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_canWinNim(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, true},
		{7, true},
		{8, false},
		{9, true},
		{10, true},
		{11, true},
		{12, false},
		{13, true},
	}

	for _, test := range tests {
		actual := canWinNim(test.n)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
