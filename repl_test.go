package main

import "testing"

func TestGetWords(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello world",
			},
		},
		{
			input: "Hello world",
			expected: []string{
				"hello world",
			},
		},
	}

	for _, test := range tests {
		words := getWords(test.input)
		if len(words) != len(test.expected) {
			t.Errorf("the lengths are not equal: %v vs %v", len(words), len(test.expected))
			continue
		}
		for i := range words {
			word := words[i]
			expected := test.expected[i]
			if word != expected {
				t.Errorf("%v does not equal %v", word, expected)
			}
		}
	}
}
