package prefixid_test

import (
	"testing"

	"github.com/jasonKoogler/prefixid"
)

func TestStringPrefixer_Prefix(t *testing.T) {
	prefixer := prefixid.StringPrefixer{}

	testCases := []struct {
		prefix   string
		id       string
		expected string
	}{
		{"usr", "123", "usr_123"},
		{"pst", "abc", "pst_abc"},
		{"", "xyz", "_xyz"},
		{"cmt", "", "cmt_"},
	}

	for _, tc := range testCases {
		t.Run(tc.prefix+"_"+tc.id, func(t *testing.T) {
			result := prefixer.Attach(tc.prefix, tc.id)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestStringPrefixer_Unprefix(t *testing.T) {
	prefixer := prefixid.StringPrefixer{}

	testCases := []struct {
		prefix     string
		prefixedID string
		expected   string
		ok         bool
	}{
		{"usr", "usr_123", "123", true},
		{"pst", "pst_abc", "abc", true},
		{"", "_xyz", "xyz", true},
		{"cmt", "cmt_", "", true},
		{"usr", "invalid_123", "", false},
		{"usr", "usr123", "", false},
		{"usr", "usrx_123", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.prefix+"_"+tc.prefixedID, func(t *testing.T) {
			result, ok := prefixer.Detach(tc.prefix, tc.prefixedID)
			if ok != tc.ok {
				t.Errorf("Expected ok=%v, got %v", tc.ok, ok)
			}

			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestStringPrefixer_Parse(t *testing.T) {
	prefixer := prefixid.StringPrefixer{}

	testCases := []struct {
		input    string
		expected string
	}{
		{"123", "123"},
		{"abc", "abc"},
		{"", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, err := prefixer.Parse(tc.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
