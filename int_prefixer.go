package prefixid

import (
	"fmt"
	"strconv"
	"strings"
)

// IntPrefixer implements IDPrefixer for int IDs
type IntPrefixer struct{}

// Prefix adds a prefix to an int ID
func (p IntPrefixer) Prefix(prefix string, id int) string {
	return fmt.Sprintf("%s_%d", prefix, id)
}

// Unprefix removes the prefix from a prefixed ID string
func (p IntPrefixer) Unprefix(prefix string, prefixedID string) (string, bool) {
	expectedPrefix := fmt.Sprintf("%s_", prefix)
	if strings.HasPrefix(prefixedID, expectedPrefix) {
		return strings.TrimPrefix(prefixedID, expectedPrefix), true
	}
	return "", false
}

// Parse parses a string into an int ID
func (p IntPrefixer) Parse(s string) (int, error) {
	return strconv.Atoi(s)
}
