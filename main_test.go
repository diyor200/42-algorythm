package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		word string
		want string
	}{
		{
			name: "first",
			word: "IceCreAm",
			want: "AceCreIm",
		},
		{
			name: "second",
			word: "leetcode",
			want: "leotcede",
		},
		{
			name: "third",
			word: "a..",
			want: "a..",
		},
		{
			name: "fourth",
			word: "ai",
			want: "ia",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseVowels(tt.word))
		})
	}

	t.Log("pass")
}
