package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		s    string
		want int
	}{
		{name: "first", s: "abccccdd", want: 7},
		{name: "second", s: "a", want: 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestPalindrome(tt.s))
		})
	}

	t.Log("pass")
}
