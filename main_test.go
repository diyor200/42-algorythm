package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		s, t string
		want byte
	}{
		{name: "first", s: "abcd", t: "abcde", want: 'e'},
		{name: "second", s: "", t: "y", want: 'y'},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findTheDifference(tt.s, tt.t))
		})
	}

	t.Log("pass")
}
