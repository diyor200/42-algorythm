package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	testCases := []struct {
		name string
		num  int
		want int
	}{
		{name: "first", num: 11, want: 3},
		{name: "second", num: 2, want: 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.HammingWeight(tt.num))
		})
	}

	t.Log("pass")
}
