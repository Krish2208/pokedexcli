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
			input: "help",
			expected: []string{
				"help",
			},
		},
		{
			input: "help me",
			expected: []string{
				"help",
				"me",
			},
		},
		{
			input: "HELP me please",
			expected: []string{
				"help",
				"me",
				"please",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Expected %d words, got %d", len(cs.expected), len(actual))
			continue
		}
		for i := range actual {
			if actual[i] != cs.expected[i] {
				t.Errorf("Expected %s, got %s", cs.expected[i], actual[i])
			}
		}
	}
}
