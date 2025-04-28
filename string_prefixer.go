package prefixid

import (
	"fmt"
	"strings"
)

// StringPrefixer implements IDPrefixer for string IDs
type StringPrefixer struct{}

var _ IDPrefixer[string] = StringPrefixer{}

// Attach attaches a prefix to a string ID
func (p StringPrefixer) Attach(prefix string, id string) string {
	return fmt.Sprintf("%s_%s", prefix, id)
}

// Detach detaches a prefix from a prefixed ID string
func (p StringPrefixer) Detach(prefix string, prefixedID string) (string, bool) {
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
