package roleverifier

import (
	"testing"
)

func TestVerify(t *testing.T) {
	tests := map[string]struct {
		path     string
		expected bool
	}{
		"multiple statements with wildcards": {
			path:     "test_input/multiple_statements_wc.json",
			expected: false,
		},
		"single statement with wildcard": {
			path:     "test_input/single_statement_wc.json",
			expected: false,
		},
		"single statement with no wildcard": {
			path:     "test_input/single_statement.json",
			expected: true,
		},
		"multiple statements with no wildcards": {
			path:     "test_input/multiple_statements.json",
			expected: true,
		},
		"not resource": {
			path:     "test_input/not_resource.json",
			expected: true,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			input, expected := test.path, test.expected
			result, err := Verify(input)
			if err != nil {
				t.Errorf("Verifier(\"%s\"), returned an unexpected error: %s", input, err)
			}
			if result != expected {
				t.Errorf("Verifier(\"%s\") = %t; expected = %t", input, result, expected)
			}
		})
	}
}
