package problem

import (
	. "github.com/chrisxue815/leetcode-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		root     string
		expected int
	}{
		{"1,1,1,#,#,1", 3},
	}

	for _, test := range tests {
		actual := maxDepth(Deserialize(test.root))
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
