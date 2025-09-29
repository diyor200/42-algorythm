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
		want []int
	}{
		{name: "first", arr: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
		{name: "second", arr: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.ProductExceptSelf(tt.arr))
		})
	}

	t.Log("pass")
}
