package prefixid

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// UUIDPrefixer implements IDPrefixer for UUID IDs
type UUIDPrefixer struct{}

var _ IDPrefixer[uuid.UUID] = UUIDPrefixer{}

// Attach attaches a prefix to a UUID ID
func (p UUIDPrefixer) Attach(prefix string, id uuid.UUID) string {
	return fmt.Sprintf("%s_%s", prefix, id.String())
}

// Detach detaches a prefix from a prefixed ID string
func (p UUIDPrefixer) Detach(prefix string, prefixedID string) (string, bool) {
	expectedPrefix := fmt.Sprintf("%s_", prefix)
	if strings.HasPrefix(prefixedID, expectedPrefix) {
		return strings.TrimPrefix(prefixedID, expectedPrefix), true
	}
	return "", false
}

// Parse parses a string into a UUID
func (p UUIDPrefixer) Parse(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}
