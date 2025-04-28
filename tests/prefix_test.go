package prefixid_test

import (
	"testing"

	"github.com/jasonKoogler/prefixid"
)

func TestNewRegistry(t *testing.T) {
	registry := prefixid.NewRegistry[string]()
	if registry == nil {
		t.Fatal("NewRegistry returned nil")
	}

	if types := registry.GetEntityTypes(); len(types) != 0 {
		t.Errorf("Expected empty registry, got %v entity types", types)
	}
}

func TestNewRegistryWithPrefixes(t *testing.T) {
	prefixMap := map[string]string{
		"user": "usr",
		"post": "pst",
	}

	registry := prefixid.NewRegistryWithPrefixes[string](prefixMap)
	if registry == nil {
		t.Fatal("NewRegistryWithPrefixes returned nil")
	}

	// Register prefixers to be able to get entity types
	registry.Register("user", "usr", prefixid.StringPrefixer{})
	registry.Register("post", "pst", prefixid.StringPrefixer{})

	types := registry.GetEntityTypes()
	if len(types) != 2 {
		t.Errorf("Expected 2 entity types, got %d: %v", len(types), types)
	}

	// Check if all expected types exist
	expectedTypes := map[string]bool{"user": false, "post": false}
	for _, entityType := range types {
		expectedTypes[entityType] = true
	}

	for entityType, found := range expectedTypes {
		if !found {
			t.Errorf("Expected entity type %s not found", entityType)
		}
	}
}

func TestRegister(t *testing.T) {
	registry := prefixid.NewRegistry[string]()

	// Register a prefixer
	registry.Register("user", "usr", prefixid.StringPrefixer{})

	types := registry.GetEntityTypes()
	if len(types) != 1 || types[0] != "user" {
		t.Errorf("Expected ['user'], got %v", types)
	}

	// Register another prefixer
	registry.Register("post", "pst", prefixid.StringPrefixer{})

	types = registry.GetEntityTypes()
	if len(types) != 2 {
		t.Errorf("Expected 2 entity types, got %d: %v", len(types), types)
	}

	// Update an existing prefixer
	registry.Register("user", "u", prefixid.StringPrefixer{})

	// Test if the prefix was updated by creating a prefixed ID
	prefixedID, err := registry.PrefixID("user", "123")
	if err != nil {
		t.Errorf("Failed to prefix ID: %v", err)
	}

	if prefixedID != "u_123" {
		t.Errorf("Expected 'u_123', got %s", prefixedID)
	}
}

func TestPrefixID(t *testing.T) {
	registry := prefixid.NewRegistry[string]()
	registry.Register("user", "usr", prefixid.StringPrefixer{})

	// Test valid prefix
	prefixedID, err := registry.PrefixID("user", "123")
	if err != nil {
		t.Errorf("Failed to prefix ID: %v", err)
	}

	if prefixedID != "usr_123" {
		t.Errorf("Expected 'usr_123', got %s", prefixedID)
	}

	// Test unregistered entity type
	_, err = registry.PrefixID("post", "456")
	if err == nil {
		t.Error("Expected error for unregistered entity type, got nil")
	}
}

func TestParsePrefixedID(t *testing.T) {
	registry := prefixid.NewRegistry[string]()
	registry.Register("user", "usr", prefixid.StringPrefixer{})

	// Test valid prefixed ID
	id, err := registry.ParsePrefixedID("user", "usr_123")
	if err != nil {
		t.Errorf("Failed to parse prefixed ID: %v", err)
	}

	if id != "123" {
		t.Errorf("Expected '123', got %s", id)
	}

	// Test invalid prefix
	_, err = registry.ParsePrefixedID("user", "invalid_123")
	if err == nil {
		t.Error("Expected error for invalid prefix, got nil")
	}

	// Test unregistered entity type
	_, err = registry.ParsePrefixedID("post", "pst_456")
	if err == nil {
		t.Error("Expected error for unregistered entity type, got nil")
	}
}

func TestMatchPrefix(t *testing.T) {
	registry := prefixid.NewRegistry[string]()
	registry.Register("user", "usr", prefixid.StringPrefixer{})
	registry.Register("post", "pst", prefixid.StringPrefixer{})

	// Test matching user prefix
	entityType, rawID, ok := registry.MatchPrefix("usr_123")
	if !ok {
		t.Error("Expected to match prefix, but did not")
	}

	if entityType != "user" {
		t.Errorf("Expected entity type 'user', got %s", entityType)
	}

	if rawID != "123" {
		t.Errorf("Expected raw ID '123', got %s", rawID)
	}

	// Test matching post prefix
	entityType, rawID, ok = registry.MatchPrefix("pst_456")
	if !ok {
		t.Error("Expected to match prefix, but did not")
	}

	if entityType != "post" {
		t.Errorf("Expected entity type 'post', got %s", entityType)
	}

	if rawID != "456" {
		t.Errorf("Expected raw ID '456', got %s", rawID)
	}

	// Test non-matching prefix
	_, _, ok = registry.MatchPrefix("invalid_789")
	if ok {
		t.Error("Expected not to match prefix, but did")
	}
}
