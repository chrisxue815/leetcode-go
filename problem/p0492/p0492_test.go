package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_constructRectangle(t *testing.T) {
	tests := []struct {
		area     int
		expected []int
	}{
		{4, []int{2, 2}},
		{15, []int{5, 3}},
		{33, []int{11, 3}},
		{41, []int{41, 1}},
	}

	for _, test := range tests {
		actual := constructRectangle(test.area)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
