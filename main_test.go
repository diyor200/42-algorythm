package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	testCases := []struct {
		name string
		arr  []int
		want bool
	}{
		{name: "first", arr: []int{1, 2, 3, 1}, want: true},
		{name: "second", arr: []int{1, 2, 3, 4}, want: false},
		{name: "third", arr: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.ContainsDuplicate(tt.arr))
		})
	}

	t.Log("pass")
}
