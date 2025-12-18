package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		n    int
		want bool
	}{
		{name: "first", n: 2, want: true},
		{name: "second", n: 3, want: false},
		{name: "third", n: 4, want: true},
		{name: "fourth", n: 5, want: false},
		{name: "fifth", n: 8, want: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, divisorGame(tt.n))
		})
	}

	t.Log("pass")
}
