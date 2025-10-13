package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	root1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}

	root2 := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 98},
		},
	}

	root3 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 1},
	}

	root4 := &TreeNode{
		Val:  0,
		Left: &TreeNode{Val: -1},
	}

	testCases := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{name: "first", root: root1, want: true},
		{name: "second", root: root2, want: false},
		{name: "third", root: root3, want: false},
		{name: "fourth", root: root4, want: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.IsValidBST(tt.root))
		})
	}

	t.Log("pass")
}
