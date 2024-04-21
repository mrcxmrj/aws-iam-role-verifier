package main

import (
	"testing"
)

func TestVerifier(t *testing.T) {
	tests := map[string]struct {
		path     string
		expected bool
	}{
		"multiple statements": {
			path:     "test_input/multiple_statements.json",
			expected: false,
		},
		"single statement": {
			path:     "test_input/single_statement.json",
			expected: false,
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
