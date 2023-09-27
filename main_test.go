package main

import (
	"testing"
)

func TestDecoder(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"LLRR=", "210122"},
		{"==RLL", "000210"},
		{"=LLRR", "221012"},
		// Add more test cases here
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := Decoder(test.input)
			if result != test.expected {
				t.Errorf("Expected: %s, Got: %s", test.expected, result)
			}
		})
	}
}
