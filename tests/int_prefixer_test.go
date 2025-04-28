package prefixid_test

import (
	"fmt"
	"testing"

	"github.com/jasonKoogler/prefixid"
)

func TestIntPrefixer_Prefix(t *testing.T) {
	prefixer := prefixid.IntPrefixer{}

	testCases := []struct {
		prefix   string
		id       int
		expected string
	}{
		{"usr", 123, "usr_123"},
		{"pst", 0, "pst_0"},
		{"", 456, "_456"},
		{"cmt", -789, "cmt_-789"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_%d", tc.prefix, tc.id), func(t *testing.T) {
			result := prefixer.Attach(tc.prefix, tc.id)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestIntPrefixer_Unprefix(t *testing.T) {
	prefixer := prefixid.IntPrefixer{}

	testCases := []struct {
		prefix     string
		prefixedID string
		expected   string
		ok         bool
	}{
		{"usr", "usr_123", "123", true},
		{"pst", "pst_0", "0", true},
		{"", "_456", "456", true},
		{"cmt", "cmt_-789", "-789", true},
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

func TestIntPrefixer_Parse(t *testing.T) {
	prefixer := prefixid.IntPrefixer{}

	testCases := []struct {
		input       string
		expected    int
		expectError bool
	}{
		{"123", 123, false},
		{"0", 0, false},
		{"-789", -789, false},
		{"abc", 0, true},
		{"", 0, true},
		{"123abc", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, err := prefixer.Parse(tc.input)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if result != tc.expected {
					t.Errorf("Expected %d, got %d", tc.expected, result)
				}
			}
		})
	}
}
