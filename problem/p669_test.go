package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
	. "leetcode-go/util"
)

func TestTrimBST(t *testing.T) {
	tests := []struct {
		root     string
		l        int
		r        int
		expected string
	}{
		{"1,0,2", 1, 2, "1,#,2"},
		{"3,0,4,#,2,#,#,1", 1, 3, "3,2,#,1"},
	}

	for _, test := range tests {
		actualTree := trimBST(Deserialize(test.root), test.l, test.r)
		actual := actualTree.Serialize()
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
