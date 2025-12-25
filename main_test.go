package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name      string
		columnNum int
		want      string
	}{
		{name: "first", columnNum: 1, want: "A"},
		{name: "second", columnNum: 28, want: "AB"},
		{name: "third", columnNum: 701, want: "ZY"},
		{name: "fourth", columnNum: 52, want: "AZ"},
		{name: "fifth", columnNum: 78, want: "BZ"},
		{name: "fifth", columnNum: 728, want: "AAZ"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, convertToTitle(tt.columnNum))
		})
	}

	t.Log("pass")
}
