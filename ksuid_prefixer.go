package prefixid

import (
	"fmt"
	"strings"

	"github.com/segmentio/ksuid"
)

// KSUIDPrefixer implements IDPrefixer for KSUID IDs
type KSUIDPrefixer struct{}

// Prefix adds a prefix to a KSUID ID
func (p KSUIDPrefixer) Prefix(prefix string, id ksuid.KSUID) string {
	return fmt.Sprintf("%s_%s", prefix, id.String())
}

// Unprefix removes the prefix from a prefixed ID string
func (p KSUIDPrefixer) Unprefix(prefix string, prefixedID string) (string, bool) {
	expectedPrefix := fmt.Sprintf("%s_", prefix)
	if strings.HasPrefix(prefixedID, expectedPrefix) {
		return strings.TrimPrefix(prefixedID, expectedPrefix), true
	}
	return "", false
}

// Parse parses a string into a KSUID
func (p KSUIDPrefixer) Parse(s string) (ksuid.KSUID, error) {
	return ksuid.Parse(s)
}
