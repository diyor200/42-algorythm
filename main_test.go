package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	testCases := []struct {
		name   string
		target int
		arr    []int
		want   []int
	}{
		{name: "first", target: 9, arr: []int{2, 7, 11, 15}, want: []int{0, 1}},
		{name: "second", target: 6, arr: []int{3, 2, 4}, want: []int{1, 2}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.TwoSum(tt.arr, tt.target))
		})
	}

	t.Log("pass")
}
