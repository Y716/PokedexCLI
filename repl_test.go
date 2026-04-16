package main

import (
	"strings"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Charmander  Bulbasaur PIKACHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "single",
			expected: []string{"single"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected length: %d, got: %d", len(c.expected), len(actual))
			t.Errorf("expected slices: %v, got: %v", strings.Join(c.expected, ","), strings.Join(actual, ","))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected word: %s, got: %s", expectedWord, word)
			}
		}
	}

}
