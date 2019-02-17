package problem

import (
	"testing"
	"github.com/stretchr/testify/assert"
	. "github.com/chrisxue815/leetcodego/util"
)

func Test_tree2str(t *testing.T) {
	tests := []struct {
		root     string
		expected string
	}{
		{"1,2,3,4", "1(2(4))(3)"},
		{"1,2,3,#,4", "1(2()(4))(3)"},
	}

	for _, test := range tests {
		actual := tree2str(Deserialize(test.root))
		assert.Equal(t, test.expected, actual, "%+v", test)
	}
}
