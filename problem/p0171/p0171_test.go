package problem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	tests := []struct {
		expected int
		s        string
	}{
		{26, "Z"},
		{27, "AA"},
		{52, "AZ"},
		{53, "BA"},
		{702, "ZZ"},
		{703, "AAA"},
	}

	for _, test := range tests {
		actual := titleToNumber(test.s)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
