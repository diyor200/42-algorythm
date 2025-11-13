package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		s    string
		k    int
		want string
	}{
		{
			name: "first",
			s:    "abcdefg",
			k:    2,
			want: "bacdfeg",
		},
		{
			name: "second",
			s:    "abcd",
			k:    2,
			want: "bacd",
		},
		{
			name: "third",
			s:    "hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl",
			k:    39,
			want: "fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseStr(tt.s, tt.k))
		})
	}

	t.Log("pass")
}
