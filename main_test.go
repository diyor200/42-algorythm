package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name  string
		image [][]int
		want  [][]int
	}{
		{
			name:  "first",
			image: [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			want:  [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			name:  "second",
			image: [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			want:  [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, flipAndInvertImage(tt.image))
		})
	}

	t.Log("pass")
}
