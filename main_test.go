package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	testCases := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "first", s: "anagram", t: "nagaram", want: true},
		{name: "second", s: "rat", t: "car", want: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.IsAnagram(tt.s, tt.t))
		})
	}

	t.Log("pass")
}
