package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	testCases := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{name: "first", root: root, want: [][]int{{3}, {9, 20}, {15, 7}}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.LevelOrder(tt.root))
		})
	}

	t.Log("pass")
}
