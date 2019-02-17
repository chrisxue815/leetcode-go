package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
	. "github.com/chrisxue815/leetcodego/util"
)

func Test_mergeTrees(t *testing.T) {
	tests := []struct {
		t1       string
		t2       string
		expected string
	}{
		{"", "", ""},
		{"1", "2", "3"},
		{"1,3,2,5", "2,1,3,#,4,#,7", "3,4,5,5,4,#,7"},
	}

	for _, test := range tests {
		actualTree := mergeTrees(Deserialize(test.t1), Deserialize(test.t2))
		actual := actualTree.Serialize()
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
