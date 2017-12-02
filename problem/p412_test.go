package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n        int
		expected []string
	}{
		{15, []string{
			"1",
			"2",
			"Fizz",
			"4",
			"Buzz",
			"Fizz",
			"7",
			"8",
			"Fizz",
			"Buzz",
			"11",
			"Fizz",
			"13",
			"14",
			"FizzBuzz",
		}},
	}

	for _, test := range tests {
		actual := fizzBuzz(test.n)
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
