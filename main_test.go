package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name     string
		rowIndex int
		want     []int
	}{
		{name: "first", rowIndex: 3, want: []int{1, 3, 3, 1}},
		{name: "second", rowIndex: 0, want: []int{1}},
		{name: "three", rowIndex: 1, want: []int{1, 1}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getRow(tt.rowIndex))
		})
	}

	t.Log("pass")
}
