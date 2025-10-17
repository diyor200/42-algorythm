package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name        string
		n           int
		edges       [][]int
		source      int
		destination int
		want        bool
	}{
		{name: "first", n: 3, source: 0, destination: 2, edges: [][]int{{0, 1}, {1, 2}, {2, 0}}, want: true},
		// {name: "second", n: 6, source: 0, destination: 5, edges: [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, want: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validPath(tt.n, tt.edges, tt.source, tt.destination))
		})
	}

	t.Log("pass")
}
