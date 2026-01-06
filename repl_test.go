package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"simple case": {
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		"leading and trailing spaces": {
			input:    "   foo bar   ",
			expected: []string{"foo", "bar"},
		},
		"multiple spaces between words": {
			input:    "go    is   fun",
			expected: []string{"go", "is", "fun"},
		},
		"empty string": {
			input:    "",
			expected: []string{},
		},
		"only spaces": {
			input:    "     ",
			expected: []string{},
		},
		"single word": {
			input:    "gopher",
			expected: []string{"gopher"},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {

			actual := cleanInput(c.input)
			if !reflect.DeepEqual(actual, c.expected) {
				t.Fatalf("input %q: expected %v, got %v", c.input, c.expected, actual)
			}
		})

	}
}
