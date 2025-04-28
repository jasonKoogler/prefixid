package prefixid_test

import (
	"sync"
	"testing"

	"github.com/jasonKoogler/prefixid"
)

func TestConcurrentRegistry(t *testing.T) {
	registry := prefixid.NewRegistry[string]()

	// Register initial entity types
	registry.Register("user", "usr", prefixid.StringPrefixer{})
	registry.Register("post", "pst", prefixid.StringPrefixer{})

	// Number of concurrent operations
	const numOps = 1000

	// Test concurrent registration
	var wg sync.WaitGroup
	wg.Add(numOps)

	for i := 0; i < numOps; i++ {
		go func(i int) {
			defer wg.Done()

			// Alternate between updating existing and adding new prefixers
			if i%2 == 0 {
				registry.Register("user", "usr", prefixid.StringPrefixer{})
			} else {
				registry.Register("comment", "cmt", prefixid.StringPrefixer{})
			}
		}(i)
	}

	wg.Wait()

	// Test concurrent prefix ID operations
	wg.Add(numOps * 2)
	errors := make(chan error, numOps*2)

	for i := 0; i < numOps; i++ {
		// Test concurrent PrefixID
		go func() {
			defer wg.Done()

			_, err := registry.PrefixID("user", "123")
			if err != nil {
				errors <- err
			}
		}()

		// Test concurrent MatchPrefix
		go func() {
			defer wg.Done()

			_, _, ok := registry.MatchPrefix("usr_123")
			if !ok {
				errors <- &matchError{message: "Failed to match prefix usr_123"}
			}
		}()
	}

	wg.Wait()
	close(errors)

	// Check for errors
	for err := range errors {
		if err != nil {
			t.Errorf("Concurrent operation error: %v", err)
		}
	}

	// Verify registry state after concurrent operations
	types := registry.GetEntityTypes()
	expectedTypes := map[string]bool{
		"user":    false,
		"post":    false,
		"comment": false,
	}

	for _, entityType := range types {
		expectedTypes[entityType] = true
	}

	for entityType, found := range expectedTypes {
		if !found {
			t.Errorf("Expected entity type %s not found after concurrent operations", entityType)
		}
	}
}

// Custom error type for match failures
type matchError struct {
	message string
}

func (e *matchError) Error() string {
	return e.message
}
