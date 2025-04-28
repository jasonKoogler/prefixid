package prefixid_test

import (
	"testing"

	"github.com/jasonKoogler/prefixid"
	"github.com/segmentio/ksuid"
)

func TestKSUIDPrefixer_Prefix(t *testing.T) {
	prefixer := prefixid.KSUIDPrefixer{}

	// Create test KSUIDs
	ksuid1, _ := ksuid.Parse("0ujtsYcgvSTl8PAuAdqWYSMnLOv")
	ksuid2, _ := ksuid.Parse("0ujzPyRiIAUAYEjrEI3fYfLRlR0")

	testCases := []struct {
		prefix   string
		id       ksuid.KSUID
		expected string
	}{
		{"usr", ksuid1, "usr_0ujtsYcgvSTl8PAuAdqWYSMnLOv"},
		{"pst", ksuid2, "pst_0ujzPyRiIAUAYEjrEI3fYfLRlR0"},
		{"", ksuid1, "_0ujtsYcgvSTl8PAuAdqWYSMnLOv"},
	}

	for _, tc := range testCases {
		t.Run(tc.prefix+"_"+tc.id.String(), func(t *testing.T) {
			result := prefixer.Attach(tc.prefix, tc.id)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestKSUIDPrefixer_Unprefix(t *testing.T) {
	prefixer := prefixid.KSUIDPrefixer{}

	testCases := []struct {
		prefix     string
		prefixedID string
		expected   string
		ok         bool
	}{
		{"usr", "usr_0ujtsYcgvSTl8PAuAdqWYSMnLOv", "0ujtsYcgvSTl8PAuAdqWYSMnLOv", true},
		{"pst", "pst_0ujzPyRiIAUAYEjrEI3fYfLRlR0", "0ujzPyRiIAUAYEjrEI3fYfLRlR0", true},
		{"", "_0ujtsYcgvSTl8PAuAdqWYSMnLOv", "0ujtsYcgvSTl8PAuAdqWYSMnLOv", true},
		{"usr", "invalid_0ujtsYcgvSTl8PAuAdqWYSMnLOv", "", false},
		{"usr", "usr0ujtsYcgvSTl8PAuAdqWYSMnLOv", "", false},
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

func TestKSUIDPrefixer_Parse(t *testing.T) {
	prefixer := prefixid.KSUIDPrefixer{}

	validKSUID := "0ujtsYcgvSTl8PAuAdqWYSMnLOv"
	expectedKSUID, _ := ksuid.Parse(validKSUID)

	testCases := []struct {
		input       string
		expected    ksuid.KSUID
		expectError bool
	}{
		{validKSUID, expectedKSUID, false},
		{"invalid-ksuid", ksuid.KSUID{}, true},
		{"", ksuid.KSUID{}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, err := prefixer.Parse(tc.input)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, but got nil for input: %s", tc.input)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if result != tc.expected {
					t.Errorf("Expected %s, got %s", tc.expected, result)
				}
			}
		})
	}
}
