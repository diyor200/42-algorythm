package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		num  int
		want bool
	}{
		{name: "first", num: 14, want: false},
		{name: "second", num: 16, want: true},
		{name: "third", num: 25, want: true},
		{name: "fourth", num: 9, want: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isPerfectSquare(tt.num))
		})
	}

	t.Log("pass")
}
