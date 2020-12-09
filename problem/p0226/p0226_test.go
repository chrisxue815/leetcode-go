package problem

import (
	"github.com/stretchr/testify/assert"
	. "github.com/chrisxue815/leetcode-go/util"
	"testing"
)

func Test_invertTree(t *testing.T) {
	tests := []struct {
		root     string
		expected string
	}{
		{"4,2,7,1,3,6,9", "4,7,2,9,6,3,1"},
		{"1,2,3,#,4,5", "1,3,2,#,5,4"},
	}

	for _, test := range tests {
		actualTree := invertTree(Deserialize(test.root))
		actual := actualTree.Serialize()
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
