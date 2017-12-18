package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
	. "leetcode-go/util"
)

func Test_convertBST(t *testing.T) {
	tests := []struct {
		root     string
		expected string
	}{
		{"5,2,13", "18,20,13"},
	}

	for _, test := range tests {
		actualTree := convertBST(Deserialize(test.root))
		actual := actualTree.Serialize()
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
