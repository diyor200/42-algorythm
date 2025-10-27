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
		{name: "first", num: 5, want: 2},
		{name: "second", num: 1, want: 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findComplement(tt.num))
		})
	}

	t.Log("pass")
}
