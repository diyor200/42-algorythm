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
		{name: "first", word: "Let's take LeetCode contest", want: "s'teL ekat edoCteeL tsetnoc"},
		{
			name: "second",
			word: "Mr Ding",
			want: "rM gniD",
		},
		{
			name: "third",
			word: "  ",
			want: "  ",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseWords(tt.word))
		})
	}

	t.Log("pass")
}
