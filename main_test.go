package main

import (
	"testing"
)

func TestVerifier(t *testing.T) {
	tests := map[string]struct {
		path     string
		expected bool
	}{
		"wrong input": {
			path:     "wrong_input",
			expected: false,
		},
		"multiple statements": {
			path:     "test_input/statements.json",
			expected: true,
		},
		"single statement": {
			path:     "test_input/test.json",
			expected: true,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			input, expected := test.path, test.expected
			result, err := verifier(input)
			if err != nil {
				t.Errorf("verifier(\"%s\"), returned an unexpected error: %s", input, err)
			}
			if result != expected {
				t.Errorf("verifier(\"%s\") = %t; expected = %t", input, result, expected)
			}
		})
	}
}
