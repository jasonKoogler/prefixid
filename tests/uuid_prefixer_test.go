package prefixid_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jasonKoogler/prefixid"
)

func TestUUIDPrefixer_Prefix(t *testing.T) {
	prefixer := prefixid.UUIDPrefixer{}

	// Create some test UUIDs
	uuid1 := uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479")
	uuid2 := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	testCases := []struct {
		prefix   string
		id       uuid.UUID
		expected string
	}{
		{"usr", uuid1, "usr_f47ac10b-58cc-0372-8567-0e02b2c3d479"},
		{"pst", uuid2, "pst_00000000-0000-0000-0000-000000000000"},
		{"", uuid1, "_f47ac10b-58cc-0372-8567-0e02b2c3d479"},
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

func TestUUIDPrefixer_Unprefix(t *testing.T) {
	prefixer := prefixid.UUIDPrefixer{}

	testCases := []struct {
		prefix     string
		prefixedID string
		expected   string
		ok         bool
	}{
		{"usr", "usr_f47ac10b-58cc-0372-8567-0e02b2c3d479", "f47ac10b-58cc-0372-8567-0e02b2c3d479", true},
		{"pst", "pst_00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000", true},
		{"", "_f47ac10b-58cc-0372-8567-0e02b2c3d479", "f47ac10b-58cc-0372-8567-0e02b2c3d479", true},
		{"usr", "invalid_f47ac10b-58cc-0372-8567-0e02b2c3d479", "", false},
		{"usr", "usrf47ac10b-58cc-0372-8567-0e02b2c3d479", "", false},
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

func TestUUIDPrefixer_Parse(t *testing.T) {
	prefixer := prefixid.UUIDPrefixer{}

	validUUID := "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	expectedUUID := uuid.MustParse(validUUID)
	zeroUUID := "00000000-0000-0000-0000-000000000000"
	expectedZeroUUID := uuid.MustParse(zeroUUID)

	testCases := []struct {
		input       string
		expected    uuid.UUID
		expectError bool
	}{
		{validUUID, expectedUUID, false},
		{zeroUUID, expectedZeroUUID, false},
		{"invalid-uuid", uuid.UUID{}, true},
		{"", uuid.UUID{}, true},
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
					t.Errorf("Expected %s, got %s", tc.expected, result)
				}
			}
		})
	}
}
