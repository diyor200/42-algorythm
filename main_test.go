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
		{name: "first", want: []int{1, 1, 0}, arr: []int{30, 60, 90}},
		{name: "second", want: []int{1, 1, 1, 0}, arr: []int{30, 40, 50, 60}},
		{name: "third", want: []int{1, 1, 4, 2, 1, 1, 0, 0}, arr: []int{73, 74, 75, 71, 69, 72, 76, 73}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.DailyTemperatures(tt.arr))
		})
	}

	t.Log("pass")
}
