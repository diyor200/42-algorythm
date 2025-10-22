package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		num  int
		want []int
	}{
		{name: "first", num: 2, want: []int{0, 1, 1}},
		{name: "second", num: 5, want: []int{0, 1, 1, 2, 1, 2}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, optimized(tt.num))
		})
	}

	t.Log("pass")
}
