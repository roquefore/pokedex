package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "HELLO",
			expected: []string{"hello"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HELLO   WORLD  ",
			expected: []string{"hello", "world"},
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
		output := cleanInput(c.input)
		if len(output) != len(c.expected) {
			t.Errorf("expected %v, got %v", c.expected, output)
		}
		for i := range output {
			if output[i] != c.expected[i] {
				t.Errorf("expected %v, got %v", c.expected, output)
			}
		}
	}
}