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
		{name: "first", num: 43261596, want: 964176192},
		{name: "second", num: 2147483644, want: 1073741822},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseBits(tt.num))
		})
	}

	t.Log("pass")
}
