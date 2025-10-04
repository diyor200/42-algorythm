package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}
	root1 := &TreeNode{Val: 2}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 3}

	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 4, Right: &TreeNode{Val: 6}, Left: &TreeNode{Val: 3}}

	testCases := []struct {
		name string
		root *TreeNode
		want int
	}{
		{name: "first", root: root1, want: 2},
		{name: "second", root: root2, want: 3},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.InvertTree(tt.root))
		})
	}

	t.Log("pass")
}
