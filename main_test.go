package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		edges [][]int
		want  int
	}{
		{name: "first", edges: [][]int{{1, 2}, {2, 3}, {4, 2}}, want: 2},
		{name: "second", edges: [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}, want: 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findCenter(tt.edges))
		})
	}

	t.Log("pass")
}
