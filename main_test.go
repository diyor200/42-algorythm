package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	testCases := []struct {
		name string
		arr  []string
		want []string
	}{
		{name: "first", want: []string{"Alaska", "Dad"}, arr: []string{"Hello", "Alaska", "Dad", "Peace"}},
		{name: "second", want: []string{}, arr: []string{"omk"}},
		{name: "third", want: []string{"adsdf", "sfd"}, arr: []string{"adsdf", "sfd"}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.FindWords(tt.arr))
		})
	}

	t.Log("pass")
}
