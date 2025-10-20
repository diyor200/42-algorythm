package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	testCases := []struct {
		name string
		word string
		want int
	}{
		{name: "first", word: "42", want: 42},
		{name: "second", word: "   -42", want: -42},
		{name: "third", word: "4193 with words", want: 4193},
		{name: "fourth", word: "-91283472332", want: -2147483648},
		{name: "fifth", word: "1337c0d3", want: 1337},
		{name: "sixth", word: "words and 987", want: 0},
		{name: "seventh", word: "+-12", want: 0},
		{name: "eights", word: "9223372036854775808", want: 2147483647},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.MyAtoi(tt.word))
		})
	}

	t.Log("pass")
}
