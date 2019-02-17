package problem

import (
	"github.com/stretchr/testify/assert"
	. "github.com/chrisxue815/leetcodego/util"
	"testing"
)

func Test_averageOfLevels(t *testing.T) {
	tests := []struct {
		root     string
		expected []float64
	}{
		{"3,9,20,#,#,15,7", []float64{3, 14.5, 11}},
	}

	for _, test := range tests {
		actual := averageOfLevels(Deserialize(test.root))
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
