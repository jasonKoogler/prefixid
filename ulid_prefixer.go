package prefixid

import (
	"fmt"
	"strings"

	"github.com/oklog/ulid/v2"
)

// ULIDPrefixer implements IDPrefixer for ULID IDs
type ULIDPrefixer struct{}

var _ IDPrefixer[ulid.ULID] = ULIDPrefixer{}

// Attach attaches a prefix to a ULID ID
func (p ULIDPrefixer) Attach(prefix string, id ulid.ULID) string {
	return fmt.Sprintf("%s_%s", prefix, id.String())
}

// Detach detaches a prefix from a prefixed ID string
func (p ULIDPrefixer) Detach(prefix string, prefixedID string) (string, bool) {
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
