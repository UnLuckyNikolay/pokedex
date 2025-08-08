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
			input:    "We're no strangers to love",
			expected: []string{"we're", "no", "strangers", "to", "love"},
		}, {
			input:    "You know the rules and so do I",
			expected: []string{"you", "know", "the", "rules", "and", "so", "do", "i"},
		}, {
			input:    "A full commitment's what I'm thinkin' of",
			expected: []string{"a", "full", "commitment's", "what", "i'm", "thinkin'", "of"},
		}, {
			input:    "You wouldn't get this from any other guy",
			expected: []string{"you", "wouldn't", "get", "this", "from", "any", "other", "guy"},
		}, {
			input:    "  Lot's of    spaces   in this   one     ",
			expected: []string{"lot's", "of", "spaces", "in", "this", "one"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%v): expected %v, actual %v", c.input, c.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v): expected %v, actual %v", c.input, c.expected, actual)
			}
		}
	}
}
