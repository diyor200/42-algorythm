package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		num  int
		want int
	}{
		{name: "first", num: 5, want: 2},
		{name: "second", num: 3, want: 2},
		{name: "third", num: 8, want: 3},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, arrangeCoins(tt.num))
		})
	}

	t.Log("pass")
}
