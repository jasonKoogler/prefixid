package prefixid

import (
	"fmt"
	"strings"
)

// StringPrefixer implements IDPrefixer for string IDs
type StringPrefixer struct{}

// Prefix adds a prefix to a string ID
func (p StringPrefixer) Prefix(prefix string, id string) string {
	return fmt.Sprintf("%s_%s", prefix, id)
}

// Unprefix removes the prefix from a prefixed ID string
func (p StringPrefixer) Unprefix(prefix string, prefixedID string) (string, bool) {
	expectedPrefix := fmt.Sprintf("%s_", prefix)
	if strings.HasPrefix(prefixedID, expectedPrefix) {
		return strings.TrimPrefix(prefixedID, expectedPrefix), true
	}
	return "", false
}

// Parse parses a string into a string ID (no-op for strings)
func (p StringPrefixer) Parse(s string) (string, error) {
	return s, nil
}
