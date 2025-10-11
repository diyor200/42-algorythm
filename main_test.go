package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 9},
		},
		Right: &TreeNode{Val: 9}}

	testCases := []struct {
		name string
		root *TreeNode
		want int
	}{
		{name: "first", root: root, want: 275},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.SumNumbers(root))
		})
	}

	t.Log("pass")
}
