package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	cases := []struct {
		name     string
		arr      []int
		expected int
	}{
		{name: "first", arr: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, expected: 49},
		{name: "second", arr: []int{1, 1}, expected: 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, s.MaxArea(c.arr))
		})
	}

	t.Log("pass")
}
