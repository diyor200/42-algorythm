package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name        string
		columnTitle string
		want        int
	}{
		{name: "first", columnTitle: "A", want: 1},
		{name: "second", columnTitle: "AA", want: 27},
		{name: "third", columnTitle: "ZY", want: 701},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, titleToNumber(tt.columnTitle))
		})
	}

	t.Log("pass")
}
