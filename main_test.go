package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		num  int
		want int
	}{
		{name: "first", num: 16, want: 4},
		{name: "second", num: 8, want: 2},
		{name: "first", num: 12, want: 3},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mySqrt(tt.num))
		})
	}

	t.Log("pass")
}
