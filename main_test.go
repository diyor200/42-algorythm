package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name    string
		letters []byte
		target  byte
		want    byte
	}{
		{name: "first", letters: []byte{'c', 'f', 'j'}, target: 'a', want: 'c'},
		{name: "second", letters: []byte{'c', 'f', 'j'}, target: 'c', want: 'f'},
		{name: "third", letters: []byte{'x', 'x', 'y', 'y'}, target: 'z', want: 'x'},
		{name: "fourth", letters: []byte{'e', 'e', 'e', 'k', 'q', 'q', 'q', 'q', 'v', 'v', 'z'}, target: 'q', want: 'v'},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, nextGreatestLetter(tt.letters, tt.target))
		})
	}

	t.Log("pass")
}
