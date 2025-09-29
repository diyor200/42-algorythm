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
		want string
	}{
		{name: "first", word: "leet**cod*e", want: "lecoe"},
		{name: "second", word: "erase*****", want: ""},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.RemoveStars(tt.word))
		})
	}

	t.Log("pass")
}
