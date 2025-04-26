package prefixid

import (
	"fmt"
	"strings"

	"github.com/oklog/ulid/v2"
)

// ULIDPrefixer implements IDPrefixer for ULID IDs
type ULIDPrefixer struct{}

// Prefix adds a prefix to a ULID ID
func (p ULIDPrefixer) Prefix(prefix string, id ulid.ULID) string {
	return fmt.Sprintf("%s_%s", prefix, id.String())
}

// Unprefix removes the prefix from a prefixed ID string
func (p ULIDPrefixer) Unprefix(prefix string, prefixedID string) (string, bool) {
	expectedPrefix := fmt.Sprintf("%s_", prefix)
	if strings.HasPrefix(prefixedID, expectedPrefix) {
		return strings.TrimPrefix(prefixedID, expectedPrefix), true
	}
	return "", false
}

// Parse parses a string into a ULID
func (p ULIDPrefixer) Parse(s string) (ulid.ULID, error) {
	return ulid.Parse(s)
}
