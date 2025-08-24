package main

import "testing"

func TestCommands(t *testing.T) {
	//
}

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   leading trailing   ",
			expected: []string{"leading", "trailing"},
		},
		{
			input:    "SoME CaPITALized",
			expected: []string{"some", "capitalized"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("less or more elements in the slice")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word doesn't match")
			}
		}
	}
}
