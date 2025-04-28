package prefixid_test

import (
	"testing"

	"github.com/jasonKoogler/prefixid"
	"github.com/oklog/ulid/v2"
)

func TestULIDPrefixer_Prefix(t *testing.T) {
	prefixer := prefixid.ULIDPrefixer{}

	// Create test ULIDs
	ulid1, _ := ulid.Parse("01F8MECHZX3TBDSZ9PT3RV4ZMH")
	ulid2, _ := ulid.Parse("01F8MED52T6YTKZYSNKJ7MX8F6")

	testCases := []struct {
		prefix   string
		id       ulid.ULID
		expected string
	}{
		{"usr", ulid1, "usr_01F8MECHZX3TBDSZ9PT3RV4ZMH"},
		{"pst", ulid2, "pst_01F8MED52T6YTKZYSNKJ7MX8F6"},
		{"", ulid1, "_01F8MECHZX3TBDSZ9PT3RV4ZMH"},
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

func TestULIDPrefixer_Unprefix(t *testing.T) {
	prefixer := prefixid.ULIDPrefixer{}

	testCases := []struct {
		prefix     string
		prefixedID string
		expected   string
		ok         bool
	}{
		{"usr", "usr_01F8MECHZX3TBDSZ9PT3RV4ZMH", "01F8MECHZX3TBDSZ9PT3RV4ZMH", true},
		{"pst", "pst_01F8MED52T6YTKZYSNKJ7MX8F6", "01F8MED52T6YTKZYSNKJ7MX8F6", true},
		{"", "_01F8MECHZX3TBDSZ9PT3RV4ZMH", "01F8MECHZX3TBDSZ9PT3RV4ZMH", true},
		{"usr", "invalid_01F8MECHZX3TBDSZ9PT3RV4ZMH", "", false},
		{"usr", "usr01F8MECHZX3TBDSZ9PT3RV4ZMH", "", false},
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

func TestULIDPrefixer_Parse(t *testing.T) {
	prefixer := prefixid.ULIDPrefixer{}

	validULID := "01F8MECHZX3TBDSZ9PT3RV4ZMH"
	expectedULID, _ := ulid.Parse(validULID)

	testCases := []struct {
		input       string
		expected    ulid.ULID
		expectError bool
	}{
		{validULID, expectedULID, false},
		{"invalid-ulid", ulid.ULID{}, true},
		{"", ulid.ULID{}, true},
		{"01f8mechzx3tbdsz9pt3rv4zmh", expectedULID, false}, // Case-insensitive
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
