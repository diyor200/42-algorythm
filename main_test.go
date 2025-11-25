package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		s    string
		c    byte
		want []int
	}{
		{
			name: "first",
			s:    "loveleetcode",
			c:    'e',
			want: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
		{
			name: "second",
			s:    "aaab",
			c:    'b',
			want: []int{3, 2, 1, 0},
		},
		{
			name: "third",
			s:    "aaba",
			c:    'b',
			want: []int{2, 1, 0, 1},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, shortestToChar(tt.s, tt.c))
		})
	}

	t.Log("pass")
}
