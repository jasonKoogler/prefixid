package prefixid

import (
	"fmt"
	"sync"
)

// Generic IDPrefixer interface
type IDPrefixer[T any] interface {
	Prefix(prefix string, id T) string
	Unprefix(prefix string, prefixedID string) (string, bool)
	Parse(s string) (T, error)
}

// Generic Registry
type Registry[T any] struct {
	prefixes  map[string]string
	prefixers map[string]IDPrefixer[T]
	mutex     sync.RWMutex
}

// NewRegistry creates a new prefix registry
func NewRegistry[T any]() *Registry[T] {
	return &Registry[T]{
		prefixes:  make(map[string]string),
		prefixers: make(map[string]IDPrefixer[T]),
	}
}

// NewRegistryWithPrefixes creates a new registry with predefined prefixes
func NewRegistryWithPrefixes[T any](prefixMap map[string]string) *Registry[T] {
	return &Registry[T]{
		prefixes:  prefixMap,
		prefixers: make(map[string]IDPrefixer[T]),
	}
}

// Register adds or updates a prefix for an entity type
func (r *Registry[T]) Register(entityType, prefix string, prefixer IDPrefixer[T]) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.prefixes[entityType] = prefix
	r.prefixers[entityType] = prefixer
}

// GetEntityTypes returns all registered entity types
func (r *Registry[T]) GetEntityTypes() []string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	types := make([]string, 0, len(r.prefixes))
	for entityType := range r.prefixes {
		types = append(types, entityType)
	}
	return types
}

// PrefixID creates a prefixed ID string for an entity type and ID
func (r *Registry[T]) PrefixID(entityType string, id T) (string, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	prefix, ok := r.prefixes[entityType]
	if !ok {
		return "", fmt.Errorf("no prefix registered for entity type: %s", entityType)
	}

	prefixer, ok := r.prefixers[entityType]
	if !ok {
		return "", fmt.Errorf("no prefixer registered for entity type: %s", entityType)
	}

	return prefixer.Prefix(prefix, id), nil
}

// ParsePrefixedID attempts to parse a prefixed ID string for a given entity type
func (r *Registry[T]) ParsePrefixedID(entityType, prefixedID string) (T, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var zero T

	prefix, ok := r.prefixes[entityType]
	if !ok {
		return zero, fmt.Errorf("no prefix registered for entity type: %s", entityType)
	}

	prefixer, ok := r.prefixers[entityType]
	if !ok {
		return zero, fmt.Errorf("no prefixer registered for entity type: %s", entityType)
	}

	rawStr, ok := prefixer.Unprefix(prefix, prefixedID)
	if !ok {
		return zero, fmt.Errorf("invalid prefix format for entity type: %s", entityType)
	}

	return prefixer.Parse(rawStr)
}

// MatchPrefix tries to determine the entity type from a prefixed ID
func (r *Registry[T]) MatchPrefix(prefixedID string) (string, string, bool) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for entityType, prefix := range r.prefixes {
		prefixer := r.prefixers[entityType]
		if rawStr, ok := prefixer.Unprefix(prefix, prefixedID); ok {
			return entityType, rawStr, true
		}
	}

	return "", "", false
}
