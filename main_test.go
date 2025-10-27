package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		num  int
		want string
	}{
		{name: "first", num: 26, want: "1a"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.)
		})
	}

	t.Log("pass")
}
