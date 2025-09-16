package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	cases := []struct {
		Name     string
		Word     string
		Expected bool
	}{
		{Word: "USA", Expected: true, Name: "true case 1"},
		{Word: "FlaG", Expected: false, Name: "wrong case 1"},
		{Word: "Name", Expected: true, Name: "true case 2"},
		{Word: "TeST", Expected: false, Name: "wrong case 2"},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			assert.Equal(t, c.Expected, s.DetectCapitalUse(c.Word))
		})
	}

	t.Log("pass!")
}
