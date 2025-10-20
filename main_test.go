package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	testCases := []struct {
		name string
		x, y int
		want int
	}{
		{name: "first", x: 1, y: 4, want: 2},
		{name: "second", x: 3, y: 1, want: 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.HammingDistance(tt.x, tt.y))
		})
	}

	t.Log("pass")
}
