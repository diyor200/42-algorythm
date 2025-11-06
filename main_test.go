package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		word string
		want bool
	}{
		{
			name: "first",
			word: "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "second",
			word: "race a car",
			want: false,
		},
		{
			name: "third",
			word: " ",
			want: true,
		},
		{
			name: "fourth",
			word: "0P",
			want: false,
		},
		{
			name: "fifth",
			word: "ab_a",
			want: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isPalindrome(tt.word))
		})
	}

	t.Log("pass")
}
