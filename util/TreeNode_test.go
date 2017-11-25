package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTreeNode_Serialize_Deserialize(t *testing.T) {
	tests := []struct {
		values string
	}{
		{"1,2,3,4,5,6,7"},
		{""},
		{"1"},
		{"1,#,2,#,3"},
		{"1,2,3,4,#,#,5"},
		{"1,2,3,#,4,5"},
	}

	for _, test := range tests {
		tree := Deserialize(test.values)
		actual := tree.Serialize()
		assert.Equal(t, test.values, actual, "values=%s", test.values)
	}
}
