package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	testCases := []struct {
		name string
		n    int
		want [][]int
	}{
		{name: "first", n: 2, want: [][]int{{1}, {1, 1}}},
		{name: "second", n: 5, want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.Generate(tt.n))
		})
	}

	t.Log("pass")
}
