package prefixid

import (
	"fmt"
	"strings"

	"github.com/segmentio/ksuid"
)

// KSUIDPrefixer implements IDPrefixer for KSUID IDs
type KSUIDPrefixer struct{}

var _ IDPrefixer[ksuid.KSUID] = KSUIDPrefixer{}

// Attach attaches a prefix to a KSUID ID
func (p KSUIDPrefixer) Attach(prefix string, id ksuid.KSUID) string {
	return fmt.Sprintf("%s_%s", prefix, id.String())
}

// Detach detaches a prefix from a prefixed ID string
func (p KSUIDPrefixer) Detach(prefix string, prefixedID string) (string, bool) {
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
