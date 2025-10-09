package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	root1 := &TreeNode{Val: 1}
	left1 := &TreeNode{Val: 2, Right: &TreeNode{Val: 4}, Left: &TreeNode{Val: 3}}
	right1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}
	root1.Left = left1
	root1.Right = right1

	root2 := &TreeNode{Val: 1}
	left2 := &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}
	right2 := &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}
	root2.Left = left2
	root2.Right = right2

	testCases := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{name: "first", root: root1, want: true},
		{name: "second", root: root2, want: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.IsSymmetric(tt.root))
		})
	}

	t.Log("pass")
}
