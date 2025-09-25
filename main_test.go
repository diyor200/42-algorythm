package main

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	cases := []struct {
		name string
		arr  []int
		want []int
	}{
		{name: "first", arr: []int{1, 2, 3, 4}, want: []int{1, 1, 1, -1}},
		{name: "second", arr: []int{10, 4, 8, 3}, want: []int{1, 1, -1, -1}},
		{name: "third", arr: []int{1, 2, 3, 4, 5}, want: []int{1, 1, 1, -1, -1}},
		{name: "four", arr: []int{0, 0, 0, 0, 0}, want: []int{0, 0, 0, 0, 0}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.LeftRightDifference(tt.arr))
		})
	}

	t.Log("pass")
}
